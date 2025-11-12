package aaa

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (u *Service) ListenSIDEncryption(ctx context.Context) error {
	eventChan := make(chan *bindings.AAASIDEncryptionRequested, 100)
	sub, err := u.contract.WatchSIDEncryptionRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		[]common.Address{u.nodeAddress},
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchSIDEncryptionRequested: %w", err)
	}

	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
	}()

	for {
		select {
		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg(models.ErrorSubscribtion.Error())
			return err

		case evt := <-eventChan:
			u.handleSIDEncryptionEvent(ctx, evt)

		case <-ctx.Done():
			u.logger.Info().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}

func (u *Service) handleSIDEncryptionEvent(ctx context.Context, evt *bindings.AAASIDEncryptionRequested) {
	if evt == nil {
		u.logger.Error().Msg("received nil event")
		return
	}

	if u.nodeAddress.Hex() != evt.Node.Hex() {
		return
	}

	b64 := base64.StdEncoding.EncodeToString(evt.Sid)
	u.logger.Debug().Msgf("SID: %s", b64)

	encryptedSID, err := PublicEncrypt(evt.Sid, evt.UserPK[:])
	if err != nil {
		return
	}

	transactOpts, err := u.newTransactor(ctx)
	if err != nil {
		return
	}

	tx, err := u.contract.SubmitEncryptedSID(
		transactOpts,
		evt.Pid,
		encryptedSID,
	)
	if err != nil {
		return
	}

	u.logger.Debug().Msgf("Submitted encrypted sid. Tx: %s", tx.Hash().Hex())
}
