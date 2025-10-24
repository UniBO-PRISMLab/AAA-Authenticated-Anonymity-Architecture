package aaa

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (u *UIP) ListenPIDEncryption(ctx context.Context) error {
	pidRequestedSig := crypto.Keccak256Hash([]byte("PIDEncryptionRequested(bytes32,address,bytes32,bytes32)"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{u.contractAddress},
		Topics:    [][]common.Hash{{pidRequestedSig}},
	}

	logs := make(chan types.Log)
	sub, err := u.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}
	defer sub.Unsubscribe()

	// transactOpts, err := u.loadTransactor(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to load transactor: %w", err)
	// }

	for {
		select {
		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg("subscription error")
			return err

		case vLog := <-logs:
			_, err := u.contract.ParsePIDEncryptionRequested(vLog)
			if err != nil {
				u.logger.Error().Err(err).Msg("failed to parse PIDEncryptionRequested event")
				continue
			}

		case <-ctx.Done():
			u.logger.Info().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}
