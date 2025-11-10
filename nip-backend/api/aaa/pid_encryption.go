package aaa

import (
	"context"
	"fmt"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u *Service) ListenPIDEncryption(ctx context.Context) error {
	eventChan := make(chan *bindings.AAAPIDEncryptionRequested, 100)
	sub, err := u.contract.WatchPIDEncryptionRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil, // TODO: filter only events for this node
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchPIDEncryptionRequested: %w", err)
	}

	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
	}()

	for {
		select {

		case evt := <-eventChan:
			go u.handlePIDEncryptionEvent(ctx, evt)

		case err := <-sub.Err():
			return err

		case <-ctx.Done():
			u.logger.Info().Msg("context cancelled, stopping watcher")
			return nil
		}
	}
}

func (u *Service) handlePIDEncryptionEvent(ctx context.Context, evt *bindings.AAAPIDEncryptionRequested) {
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
		Msg("Received PIDEncryptionRequested")

	encryptedPID, err := SymEncrypt(evt.Pid[:], evt.SymK[:])
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorPIDEncryption.Error())
		return
	}

	transactOpts, err := u.newTransactor(ctx)
	if err != nil {
		return
	}

	tx, err := u.contract.SubmitEncryptedPID(
		transactOpts,
		evt.Pid,
		evt.Sid,
		encryptedPID,
	)
	if err != nil {
		u.logger.Error().Err(err).Msg("failed to submit encrypted PID")
		return
	}

	u.logger.Debug().Msgf("Submitted encrypted PID. Tx: %s", tx.Hash().Hex())
}
