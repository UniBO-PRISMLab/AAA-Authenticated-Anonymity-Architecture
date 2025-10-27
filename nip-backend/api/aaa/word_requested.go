package aaa

import (
	"context"
	"strings"

	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (u *Service) ListenWordRequested(ctx context.Context) error {
	wordRequestedSig := crypto.Keccak256Hash([]byte("WordRequested(bytes32,address,bytes)"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{u.contractAddress},
		Topics:    [][]common.Hash{{wordRequestedSig}},
	}

	logs := make(chan types.Log)
	sub, err := u.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return models.ErrorSubscribeToLogs
	}
	defer sub.Unsubscribe()

	transactOpts, err := u.loadTransactor(ctx)
	if err != nil {
		return models.ErrorLoadTransactor
	}

	for {
		select {
		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg(models.ErrorSubscribtion.Error())
			return err

		case vLog := <-logs:
			event, err := u.contract.ParseWordRequested(vLog)
			if err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorParseWordRequested.Error())
				continue
			}

			if u.nodeAddress.Hex() != event.Node.Hex() {
				continue
			}

			word := strings.ToLower(u.babbler.Babble())
			encryptedWord, err := PublicEncrypt([]byte(word), event.UserPK[:])
			if err != nil {
				return models.ErrorWordEncryption
			}

			tx, err := u.contract.SubmitEncryptedWord(
				transactOpts,
				event.Pid,
				encryptedWord,
				[]byte(u.configuration.KeyPair.PublicKey),
			)
			if err != nil {
				return models.ErrorWordSubmission
			}

			u.logger.Debug().Msgf("Submitted encrypted word. Tx: %s", tx.Hash().Hex())

		case <-ctx.Done():
			u.logger.Debug().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}
