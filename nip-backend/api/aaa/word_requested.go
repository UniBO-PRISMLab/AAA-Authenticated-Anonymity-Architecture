package aaa

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u *Service) ListenWordRequested(ctx context.Context) error {
	eventChan := make(chan *bindings.AAAContractWordRequested, 100)
	sub, err := u.contract.WatchWordRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchWordRequested: %w", err)
	}

	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
	}()

	for {
		select {
		case evt := <-eventChan:
			go u.handleWordEncryptionEvent(ctx, evt)

		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg("subscription error in WordRequested listener")
			return err

		case <-ctx.Done():
			u.logger.Debug().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}

func (u *Service) handleWordEncryptionEvent(ctx context.Context, evt *bindings.AAAContractWordRequested) {
	if evt == nil {
		u.logger.Error().Msg("received nil event")
		return
	}

	if u.nodeAddress.Hex() != evt.Node.Hex() {
		u.logger.Error().Msg("not the one, skipping")
		return
	}

	u.logger.Debug().
		Str("pid", fmt.Sprintf("%x", evt.Pid)).
		Msg("Received WordRequested")

	pidB64 := base64.StdEncoding.EncodeToString(evt.Pid[:])
	if _, err := u.identityService.GetUserByPID(ctx, &pidB64); err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorUserWithPIDNotFound.Error())
		return
	}

	word := strings.ToLower(u.babbler.Babble())
	encryptedWord, err := PublicEncrypt([]byte(word), evt.UserPK[:])
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorWordEncryption.Error())
		return
	}

	transactOpts, err := u.newTransactor(ctx)
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorLoadTransactor.Error())
		return
	}

	tx, err := u.contract.SubmitEncryptedWord(
		transactOpts,
		evt.Pid,
		encryptedWord,
		[]byte(u.configuration.KeyPair.PublicKey),
	)
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorWordSubmission.Error())
		return
	}

	u.logger.Debug().Msgf("Submitted encrypted word. Tx: %s", tx.Hash().Hex())
}
