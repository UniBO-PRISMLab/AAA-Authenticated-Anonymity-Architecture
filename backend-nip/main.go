package main

import (
	"context"
	"os"

	"github.com/UniBO-PRISMLab/nip/api"
	"github.com/UniBO-PRISMLab/nip/api/auth"
	"github.com/UniBO-PRISMLab/nip/api/identity"
	"github.com/UniBO-PRISMLab/nip/db"
	_ "github.com/UniBO-PRISMLab/nip/docs"
	"github.com/UniBO-PRISMLab/nip/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

//	@title			National Identity Provider (NIP)
//	@version		0.0.1
//	@description	API specification for the Authenticated Anonymity Architecture (AAA), a blockchain-based solution providing robust and ethical authenticated anonymous identities with deanonymization capabilities for criminal activity cases.
//	@contact.email	m.dinelli@unibo.it
//	@contact.name	Michele Dinelli

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8888
//	@schemes	http https
//	@BasePath	/

func main() {
	var err error

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	configuration := models.NewConfiguration()

	dbpool, err := pgxpool.New(context.Background(), configuration.DB.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create connection pool")
		os.Exit(1)
	}

	defer dbpool.Close()

	db := &db.DB{
		Pool: dbpool,
	}

	repos := db.InitRepositories()

	identityService := identity.NewService(configuration, repos.Identity)
	authService := auth.NewService(configuration, repos.Auth, identityService)

	if err = api.NewServer(
		configuration,
		identityService,
		authService,
	).Listen(); err != nil {
		ctx.Done()
		os.Exit(1)
	}
}
