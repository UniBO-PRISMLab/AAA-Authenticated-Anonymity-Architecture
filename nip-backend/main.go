package main

import (
	"context"
	"os"

	"github.com/UniBO-PRISMLab/nip-backend/api"
	"github.com/UniBO-PRISMLab/nip-backend/api/aaa"
	"github.com/UniBO-PRISMLab/nip-backend/api/auth"
	"github.com/UniBO-PRISMLab/nip-backend/api/identity"
	"github.com/UniBO-PRISMLab/nip-backend/db"
	_ "github.com/UniBO-PRISMLab/nip-backend/docs"
	"github.com/UniBO-PRISMLab/nip-backend/models"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

//	@title			National Identity Provider (NIP)
//	@version		0.0.1
//	@description	API specification for the NIP server of the Authenticated Anonymity Architecture (AAA).
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
		log.Fatal().Err(err).Msg(models.ErrorUnableToCreateConnPool.Error())
		os.Exit(1)
	}

	defer dbpool.Close()

	db := &db.DB{
		Pool: dbpool,
	}

	repos := db.InitRepositories()
	identityService := identity.NewService(configuration, repos.Identity)

	ethClient, err := eth.DialContext(ctx, configuration.Blockchain.EthNodeURL)
	if err != nil {
		log.Error().Err(err).Msg(models.ErrorUnableToConnectToEthClient.Error())
		os.Exit(1)
	}
	defer ethClient.Close()

	uip, err := aaa.NewAAAService(
		ethClient,
		identityService,
		configuration.Blockchain.ContractAddress,
		configuration,
	)
	if err != nil {
		log.Error().Err(err).Msg(models.ErrorUnableToCreateUIPListener.Error())
	}

	authService := auth.NewService(configuration, repos.Auth, identityService, uip)

	go uip.Start(ctx)

	if err = api.NewServer(
		configuration,
		identityService,
		authService,
	).Listen(); err != nil {
		ctx.Done()
		os.Exit(1)
	}
}
