package aaa

import (
	"context"
	"fmt"
	"time"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u *Service) ListenPIDEncryption(ctx context.Context) error {
	eventChan := make(chan *bindings.AAAContractPIDEncryptionRequested)
	sub, err := u.contract.WatchPIDEncryptionRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil, // TODO: filter only events for this node
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchPIDEncryptionRequested: %w", err)
	}
	defer sub.Unsubscribe()

	for {
		select {

		case evt := <-eventChan:
			if evt == nil {
				u.logger.Error().Msg("received nil event")
				continue
			}

			if u.nodeAddress.Hex() != evt.Node.Hex() {
				continue
			}

			u.logger.Debug().
				Str("pid", fmt.Sprintf("%x", evt.Pid)).
				Str("sid", fmt.Sprintf("%x", evt.Sid)).
				Msg("Received PIDEncryptionRequested")

			encryptedPID, err := SymEncrypt(evt.Pid[:], evt.SymK[:])
			if err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorPIDEncryption.Error())
				return models.ErrorPIDEncryption
			}

			transactOpts, err := u.newTransactor(ctx)
			if err != nil {
				return models.ErrorLoadTransactor
			}

			tx, err := u.contract.SubmitEncryptedPID(
				transactOpts,
				evt.Pid,
				evt.Sid,
				encryptedPID,
			)
			if err != nil {
				u.logger.Error().Err(err).Msg("failed to submit encrypted PID")
				continue
			}

			u.logger.Debug().Msgf("Submitted encrypted PID. Tx: %s", tx.Hash().Hex())

		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg("subscription error, restarting watcher")

			time.Sleep(2 * time.Second)
			go func() {
				if restartErr := u.ListenPIDEncryption(ctx); restartErr != nil {
					u.logger.Error().Err(restartErr).Msg("failed to restart listener")
				}
			}()
			return err

		case <-ctx.Done():
			u.logger.Info().Msg("context cancelled, stopping watcher")
			return nil
		}
	}
}
