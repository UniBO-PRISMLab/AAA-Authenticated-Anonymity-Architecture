package db

import (
	"context"
	_ "embed"

	"github.com/UniBO-PRISMLab/nip/models"
)

//go:embed sql/auth/insert_pid.sql
var insertPIDQuery string

type IdentityRepository struct {
	DB *DB
}

func NewIdentityRepository(db *DB) *IdentityRepository {
	return &IdentityRepository{
		DB: db,
	}
}

func (r *IdentityRepository) IssuePID(ctx context.Context, publicKey string, pid string) (*models.PIDResponseModel, error) {
	tx, err := r.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var insertedPID string
	var insertedPublicKey string

	err = tx.QueryRow(ctx, insertPIDQuery, pid, publicKey).Scan(&insertedPID, &insertedPublicKey)
	if err != nil {
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
