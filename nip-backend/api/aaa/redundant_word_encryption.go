package aaa

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/UniBO-PRISMLab/nip-backend/api/aaa/bindings"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (u *Service) ListenRedundantWordRequested(ctx context.Context) error {
	eventChan := make(chan *bindings.AAARedundantWordRequested, 100)
	sub, err := u.contract.WatchRedundantWordRequested(
		&bind.WatchOpts{Context: ctx},
		eventChan,
		nil,
		nil,
		[]common.Address{u.nodeAddress},
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe via WatchRedundantWordRequested: %w", err)
	}

	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
	}()

	for {
		select {
		case evt := <-eventChan:
			u.handleRedundantWordEvent(ctx, evt)

		case err := <-sub.Err():
			u.logger.Error().Err(err).Msg("subscription error in RedundantWordRequested listener")
			return err

		case <-ctx.Done():
			u.logger.Debug().Msg("context cancelled, stopping listener")
			return nil
		}
	}
}

func (u *Service) handleRedundantWordEvent(ctx context.Context, evt *bindings.AAARedundantWordRequested) {
	if evt == nil {
		u.logger.Error().Msg("received nil event")
		return
	}

	if u.nodeAddress.Hex() != evt.ToNode.Hex() {
		return
	}

	pidB64 := base64.StdEncoding.EncodeToString(evt.Pid[:])
	if _, err := u.identityService.GetUserByPID(ctx, &pidB64); err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorUserWithPIDNotFound.Error())
		return
	}

	pkBytes, err := base64.StdEncoding.DecodeString(u.configuration.KeyPair.PublicKey)
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorPublicKeyDecoding.Error())
		return
	}

	pemBlock, _ := pem.Decode(pkBytes)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		u.logger.Error().Err(err).Msg(models.ErrorInvalidPublicKeyHeader.Error())
		return
	}

	encWord, err := PublicEncrypt(evt.HashedWord[:], pkBytes)
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorRedundantWordEncryption.Error())
	}

	transactOpts, err := u.newTransactor(ctx)
	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorLoadTransactor.Error())
		return
	}

	tx, err := u.contract.SubmitRedundantWord(
		transactOpts,
		evt.Pid,
		encWord,
		evt.Index,
		pkBytes,
	)

	if err != nil {
		u.logger.Error().Err(err).Msg(models.ErrorWordSubmission.Error())
		return
	}

	u.logger.Debug().Msgf("Submitted redundant word. Tx: %s", tx.Hash().Hex())
}
