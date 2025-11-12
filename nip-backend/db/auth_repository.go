package db

import (
	"context"
	_ "embed"
	"time"

	"github.com/UniBO-PRISMLab/nip-backend/models"
	"github.com/jackc/pgx/v5"
)

//go:embed sql/auth/insert_pac.sql
var insertPACQuery string

//go:embed sql/auth/get_active_pac.sql
var getActivePACQuery string

//go:embed sql/auth/delete_pac.sql
var deletePACQuery string

//go:embed sql/auth/insert_sac.sql
var insertSACQuery string

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

func (r *AuthRepository) VerifyPAC(ctx context.Context, pid *string, pac int64) (*models.PACVerificationResponseModel, error) {
	var expiration time.Time
	var foundPAC int64

	tx, err := r.DB.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, getActivePACQuery, pid, pac).Scan(&foundPAC, &expiration)
	if err != nil || foundPAC != pac {
		return nil, models.ErrorPACNotValid
	}

	_, err = tx.Exec(ctx, deletePACQuery, pid, pac)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &models.PACVerificationResponseModel{
		Valid:      foundPAC == pac,
		Expiration: expiration,
	}, nil
}

func (r *AuthRepository) IssueSAC(ctx context.Context,
	sac string,
	sid *string,
	expiration time.Time) (*models.SACResponseModel, error) {
	var err error
	var insertedSAC string
	var insertedSID string

	tx, err := r.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(ctx, insertSACQuery, sac, expiration, sid).Scan(&insertedSAC, &expiration, &insertedSID)
	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	return &models.SACResponseModel{
		SAC:        insertedSAC,
		Expiration: expiration,
	}, nil
}
