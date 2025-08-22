package db

type DB struct {
}

type Repositories struct {
	Auth     *AuthRepository
	Identity *IdentityRepository
}

func InitRepositories() *Repositories {
	return &Repositories{
		Auth:     NewAuthRepository(),
		Identity: NewIdentityRepository(),
	}
}
