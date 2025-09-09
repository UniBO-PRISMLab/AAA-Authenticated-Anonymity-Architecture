package db

import (
	"context"
	_ "embed"
	"time"

	"github.com/UniBO-PRISMLab/nip/models"
)

//go:embed sql/auth/insert_pac.sql
var insertPACQuery string

//go:embed sql/auth/get_active_pac.sql
var getActivePACQuery string

type AuthRepository struct {
	DB *DB
}

func NewAuthRepository(DB *DB) *AuthRepository {
	return &AuthRepository{
		DB: DB,
	}
}

func (r *AuthRepository) IssuePAC(ctx context.Context, pid *string, pac int64, expiration time.Time) (*models.PACResponseModel, error) {
	var insertedPAC int64
	var insertedExpiration time.Time
	var insertedPID string

	tx, err := r.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, insertPACQuery, pac, expiration, pid).Scan(&insertedPAC, &insertedExpiration, &insertedPID)
	if err != nil {
		if r.DB.IsUniqueConstraintError(err) {
			return nil, models.ErrorPKAlreadyAssociated
		}
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &models.PACResponseModel{
		PAC:        insertedPAC,
		Expiration: insertedExpiration,
	}, nil
}

func (r *AuthRepository) GetSAC(ctx context.Context) (*models.SACResponseModel, error) {
	return &models.SACResponseModel{}, nil
}

func (r *AuthRepository) VerifyPAC(ctx context.Context, pid *string, pac int64) (*models.PACVerificationResponseModel, error) {
	var expiration time.Time
	foundPAC := 0

	err := r.DB.Pool.QueryRow(ctx, getActivePACQuery, pid, pac).Scan(&foundPAC, &expiration)
	if err != nil {
		return nil, models.ErrorPACNotValid
	}

	return &models.PACVerificationResponseModel{
		Valid:      foundPAC != 0,
		Expiration: expiration,
	}, nil
}
