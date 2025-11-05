package aaa

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u *Service) ListenWordRequested(ctx context.Context) error {
	eventChan := make(chan *bindings.AAAContractWordRequested)
	sub, err := u.contract.WatchWordRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil, // TODO: filter only events for this node
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchWordRequested: %w", err)
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
				Msg("Received WordRequested")

			pidB64 := base64.StdEncoding.EncodeToString(evt.Pid[:])
			if _, err = u.identityService.GetUserByPID(ctx, &pidB64); err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorUserWithPIDNotFound.Error())
				continue
			}

			word := strings.ToLower(u.babbler.Babble())
			encryptedWord, err := PublicEncrypt([]byte(word), evt.UserPK[:])
			if err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorWordEncryption.Error())
				return models.ErrorWordEncryption
			}

			transactOpts, err := u.newTransactor(ctx)
			if err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorLoadTransactor.Error())
				return models.ErrorLoadTransactor
			}

			tx, err := u.contract.SubmitEncryptedWord(
				transactOpts,
				evt.Pid,
				encryptedWord,
				[]byte(u.configuration.KeyPair.PublicKey),
			)
			if err != nil {
				u.logger.Error().Err(err).Msg(models.ErrorWordSubmission.Error())
				return models.ErrorWordSubmission
			}

			u.logger.Debug().Msgf("Submitted encrypted word. Tx: %s", tx.Hash().Hex())

		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg("subscription error, restarting watcher")

			time.Sleep(2 * time.Second)
			go func() {
				if restartErr := u.ListenWordRequested(ctx); restartErr != nil {
					u.logger.Error().Err(restartErr).Msg("failed to restart listener")
				}
			}()
			return err

		case <-ctx.Done():
			u.logger.Debug().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}
