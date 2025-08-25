package db

import (
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
