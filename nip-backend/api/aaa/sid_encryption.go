package aaa

import (
	"context"
	"fmt"

	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (u *UIP) ListenSIDEncryption(ctx context.Context) error {
	sidRequestedSig := crypto.Keccak256Hash([]byte("SIDEncryptionRequested(bytes32,address,bytes,bytes)"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{u.contractAddress},
		Topics:    [][]common.Hash{{sidRequestedSig}},
	}

	logs := make(chan types.Log)
	sub, err := u.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}
	defer sub.Unsubscribe()

	transactOpts, err := u.loadTransactor(ctx)
	if err != nil {
		return fmt.Errorf("failed to load transactor: %w", err)
	}

	for {
		select {
		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg(models.ErrorSubscribtion.Error())
			return err

		case vLog := <-logs:
			event, err := u.contract.ParseSIDEncryptionRequested(vLog)
			if err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorParseSIDEncryptionRequested.Error())
				continue
			}

			if u.nodeAddress.Hex() != event.Node.Hex() {
				continue
			}

			encryptedSID, err := PublicEncrypt(event.Sid, event.UserPK[:])
			if err != nil {
				return models.ErrorWordEncryption
			}

			tx, err := u.contract.SubmitEncryptedSID(
				transactOpts,
				event.Pid,
				encryptedSID,
			)
			if err != nil {
				return models.ErrorWordSubmission
			}

			u.logger.Debug().Msgf("Submitted encrypted sid. Tx: %s", tx.Hash().Hex())

		case <-ctx.Done():
			u.logger.Info().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}
