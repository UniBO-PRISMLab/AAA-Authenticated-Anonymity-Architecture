package aaa

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/UniBO-PRISMLab/nip-backend/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/tjarratt/babble"
)

type Service struct {
	client          *ethclient.Client
	contract        *bindings.AAAContract
	babbler         *babble.Babbler
	configuration   models.Configuration
	contractAddress common.Address
	nodeAddress     common.Address
	logger          *zerolog.Logger
}

func NewAAAService(
	client *ethclient.Client,
	contractAddr string,
	configuration models.Configuration,
) (*Service, error) {
	logger := utils.InitServiceAdvancedLogger("AAALogger")
	addr := common.HexToAddress(contractAddr)
	nodeAddr, err := getBackendAddress(configuration.Blockchain.BlockchainPrivateKey)
	logger.Info().Msgf("UIP listener started at %s", nodeAddr.Hex())

	if err != nil {
		return nil, fmt.Errorf("failed to get backend address: %w", err)
	}

	contract, err := bindings.NewAAAContract(addr, client)
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
	}, nil
}

func (u *Service) Start(ctx context.Context) {
	go u.ListenWordRequested(ctx)
	go u.ListenPIDEncryption(ctx)
	go u.ListenSIDEncryption(ctx)
}

func (u *Service) loadTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	keyHex := u.configuration.Blockchain.BlockchainPrivateKey
	if len(keyHex) >= 2 && keyHex[:2] == "0x" {
		keyHex = keyHex[2:]
	}

	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key hex: %w", err)
	}

	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	chainID, err := u.client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %w", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	transactOpts.GasLimit = 5_000_000
	return transactOpts, nil
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
	// TODO: implement symmetric encryption
	return data, nil
}

func (u *Service) GetSIDRecord(ctx context.Context, sidBase64 string) ([]byte, []byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(sidBase64)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid base64 SID: %w", err)
	}

	sidHex := strings.TrimSpace(string(decoded))
	sidHex = strings.TrimPrefix(sidHex, "0x")
	sidBytes, err := hex.DecodeString(sidHex)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid SID hex: %w", err)
	}
	if len(sidBytes) != 32 {
		return nil, nil, fmt.Errorf("SID must be 32 bytes, got %d", len(sidBytes))
	}

	var sid [32]byte
	copy(sid[:], sidBytes)

	encPID, pk, err := u.contract.GetSIDRecord(&bind.CallOpts{Context: ctx}, sid)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read SID record: %w", err)
	}

	return encPID, pk, nil
}

func getBackendAddress(hexKey string) (common.Address, error) {
	if len(hexKey) >= 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}

	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to decode private key: %w", err)
	}

	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid private key: %w", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	return address, nil
}
