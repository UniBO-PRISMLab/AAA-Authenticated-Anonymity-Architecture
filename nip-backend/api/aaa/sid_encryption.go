package aaa

import (
	"context"
	"fmt"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u *Service) ListenSIDEncryption(ctx context.Context) error {
	eventChan := make(chan *bindings.AAASIDEncryptionRequested, 100)
	sub, err := u.contract.WatchSIDEncryptionRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil, // TODO: filter only events for this node
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
			go u.handleSIDEncryptionEvent(ctx, evt)

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

	u.logger.Debug().
		Str("pid", fmt.Sprintf("%x", evt.Pid)).
		Str("sid", fmt.Sprintf("%x", evt.Sid)).
		Msg("Received SIDEncryptionRequested")

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
