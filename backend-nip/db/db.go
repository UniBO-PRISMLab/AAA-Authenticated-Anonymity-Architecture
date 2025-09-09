package db

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

type Repositories struct {
	Auth     *AuthRepository
	Identity *IdentityRepository
}

func (db *DB) InitRepositories() *Repositories {
	return &Repositories{
		Auth:     NewAuthRepository(db),
		Identity: NewIdentityRepository(db),
	}
}

func (db *DB) IsUniqueConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return true
		}
	}
	return false
}
