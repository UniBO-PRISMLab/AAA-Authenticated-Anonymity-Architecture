package aaa

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/api/identity"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/UniBO-PRISMLab/nip-backend/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/tjarratt/babble"
)

type Service struct {
	client          *ethclient.Client
	contract        *bindings.AAA
	babbler         *babble.Babbler
	configuration   models.Configuration
	contractAddress common.Address
	nodeAddress     common.Address
	logger          *zerolog.Logger
	identityService *identity.Service
}

func NewAAAService(
	client *ethclient.Client,
	identityService *identity.Service,
	contractAddr string,
	configuration models.Configuration,
) (*Service, error) {
	logger := utils.InitServiceAdvancedLogger("AAALogger")
	addr := common.HexToAddress(contractAddr)
	nodeAddr := common.HexToAddress(configuration.Blockchain.BlockchainAddress)

	contract, err := bindings.NewAAA(addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate contract: %w", err)
	}

	babbler := babble.NewBabbler()
	babbler.Count = 1

	return &Service{
		client:          client,
		contract:        contract,
		contractAddress: addr,
		nodeAddress:     nodeAddr,
		babbler:         &babbler,
		configuration:   configuration,
		logger:          logger,
		identityService: identityService,
	}, nil
}

func (u *Service) Start(ctx context.Context) {
	go u.watchLoop(ctx, "WordRequested", u.ListenWordRequested)
	go u.watchLoop(ctx, "PIDEncryption", u.ListenPIDEncryption)
	go u.watchLoop(ctx, "SIDEncryption", u.ListenSIDEncryption)
	// go u.watchLoop(ctx, "RedundantWordRequested", u.ListenRedundantWordRequested)
}

func (u *Service) watchLoop(ctx context.Context, name string, fn func(context.Context) error) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					u.logger.Error().Msgf("[%s] panic recovered: %v", name, r)
				}
			}()

			u.logger.Info().Msgf("[%s] listener started", name)
			err := fn(ctx)
			if err != nil {
				u.logger.Error().Err(err).Msgf("[%s] listener failed", name)
			}
		}()

		select {
		case <-ctx.Done():
			u.logger.Info().Msgf("[%s] listener stopped due to context", name)
			return
		case <-time.After(5 * time.Second):
			u.logger.Warn().Msgf("[%s] restarting listener after 5s", name)
		}
	}
}

func (u *Service) newTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	keyHex := strings.TrimPrefix(u.configuration.Blockchain.BlockchainPrivateKey, "0x")
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	chainID, err := u.client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	opts.GasLimit = 5_000_000

	return opts, nil
}

func PublicEncrypt(data []byte, key []byte) ([]byte, error) {
	rng := rand.Reader
	publicKeyPemBlock, _ := pem.Decode(key)
	if publicKeyPemBlock == nil || publicKeyPemBlock.Type != "PUBLIC KEY" {
		return nil, models.ErrorInvalidPublicKeyHeader
	}

	pk, err := x509.ParsePKIXPublicKey(publicKeyPemBlock.Bytes)
	if err != nil {
		return nil, models.ErrorInvalidPublicKey
	}
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rng, pk.(*rsa.PublicKey), data, nil)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

func SymEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, models.ErrorInvalidSymK
	}
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)
	enc := cipher.NewCBCEncrypter(block, iv)
	plaintext := pad(data, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	enc.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func pad(data []byte, blockSize int) []byte {
	n := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(n)}, n)
	return append(data, padding...)
}

func (u *Service) GetSIDRecord(ctx context.Context, sidBase64 string) ([]byte, []byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(sidBase64)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid base64 SID: %w", err)
	}

	if len(decoded) != 32 {
		return nil, nil, fmt.Errorf("SID must be 32 bytes, got %d", len(decoded))
	}

	var sid [32]byte
	copy(sid[:], decoded)

	encPID, pk, err := u.contract.GetSIDRecord(&bind.CallOpts{Context: ctx}, sid)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read SID record: %w", err)
	}

	return encPID, pk, nil
}

func (u *Service) SubmitSAC(ctx context.Context, sac []byte) error {
	var opts *bind.TransactOpts
	var err error
	var tx *types.Transaction

	if opts, err = u.newTransactor(ctx); err != nil {
		return models.ErrorLoadTransactor
	}
	if tx, err = u.contract.SubmitSAC(opts, sac); err != nil {
		return models.ErrorSACSubmission
	}

	u.logger.Debug().Msgf("Submitted sac %s. Tx: %s",
		base64.StdEncoding.EncodeToString(sac),
		tx.Hash().Hex(),
	)

	return nil
}
