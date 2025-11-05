package aaa

import (
	"context"
	"fmt"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u *Service) ListenSIDEncryption(ctx context.Context) error {
	eventChan := make(chan *bindings.AAAContractSIDEncryptionRequested)
	sub, err := u.contract.WatchSIDEncryptionRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil, // TODO: filter only events for this node
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchSIDEncryptionRequested: %w", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg(models.ErrorSubscribtion.Error())
			return err

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
				Msg("Received SIDEncryptionRequested")

			encryptedSID, err := PublicEncrypt(evt.Sid, evt.UserPK[:])
			if err != nil {
				return models.ErrorWordEncryption
			}

			transactOpts, err := u.newTransactor(ctx)
			if err != nil {
				return models.ErrorLoadTransactor
			}

			tx, err := u.contract.SubmitEncryptedSID(
				transactOpts,
				evt.Pid,
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
