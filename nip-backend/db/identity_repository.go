package db

import (
	"context"
	_ "embed"

	"github.com/UniBO-PRISMLab/nip-backend/models"
)

//go:embed sql/identity/insert_pid.sql
var insertPIDQuery string

//go:embed sql/identity/get_by_pid.sql
var getByPIDQuery string

type IdentityRepository struct {
	DB *DB
}

func NewIdentityRepository(db *DB) *IdentityRepository {
	return &IdentityRepository{
		DB: db,
	}
}

func (r *IdentityRepository) IssuePID(
	ctx context.Context,
	publicKey string,
	pid string,
	nonce string) (*models.PIDResponseModel, error) {
	tx, err := r.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var insertedPID string
	var insertedPublicKey string
	var insertedNonce string

	err = tx.QueryRow(ctx, insertPIDQuery, pid, publicKey, nonce).Scan(&insertedPID, &insertedPublicKey, &insertedNonce)
	if err != nil {
		if r.DB.IsUniqueConstraintError(err) {
			return nil, models.ErrorPKAlreadyAssociated
		}
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &models.PIDResponseModel{
		PID:     insertedPID,
		Message: models.MsgPIDCreated,
	}, nil
}

func (r *IdentityRepository) GetUserByPID(ctx context.Context, PID *string) (*models.User, error) {
	var pid string
	var publicKey string
	var nonce string

	err := r.DB.Pool.QueryRow(ctx, getByPIDQuery, &PID).Scan(&pid, &publicKey, &nonce)
	if err != nil {
		return nil, err
	}

	return &models.User{
		PID:       pid,
		PublicKey: publicKey,
		Nonce:     nonce,
	}, nil
}
