package models

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentType string

const (
	Development EnvironmentType = "development"
	Production  EnvironmentType = "production"
)

type DBConfiguration struct {
	DatabaseURL string
}

type BlockchainConfiguration struct {
	EthNodeURL           string
	ContractAddress      string
	BlockchainPrivateKey string
}

type Configuration struct {
	Environment EnvironmentType
	HTTPHost    string
	HTTPPort    int
	CORS        CORSConfiguration
	DB          DBConfiguration
	SK          []byte
	PublicKey   string
	Blockchain  BlockchainConfiguration
}

type CORSConfiguration struct {
	AllowOrigins []string
	AllowHeaders []string
}

func NewConfiguration() Configuration {
	var env EnvironmentType
	var err error

	if err = godotenv.Load(); err != nil {
		log.Println("Couldn't load .env file")
	}

	os.Setenv("GIN_MODE", stringFromEnv("GIN_MODE", "development"))
	os.Setenv("HTTP_HOST", "0.0.0.0")
	os.Setenv("HTTP_PORT", "8888")

	httpHost := stringOrPanic("HTTP_HOST")
	httpPort := intOrPanic("HTTP_PORT")

	databaseURL := stringOrPanic("DATABASE_URL")

	ethNodeUrl := stringOrPanic("ETH_NODE_URL")

	secretKey := [32]byte{}
	copy(secretKey[:], stringOrPanic("SK"))

	publicKey := stringOrPanic("PUBLIC_KEY")

	contractAddress := stringOrPanic("CONTRACT_ADDRESS")
	blockchainPrivateKey := stringOrPanic("BLOCKCHAIN_PRIVATE_KEY")

	if stringFromEnv("GIN_MODE", "development") == "production" {
		env = Production
	} else {
		env = Development
	}

	return Configuration{
		Environment: env,
		HTTPHost:    httpHost,
		HTTPPort:    httpPort,
		CORS: CORSConfiguration{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{},
		},
		DB: DBConfiguration{
			DatabaseURL: databaseURL,
		},
		SK:        secretKey[:],
		PublicKey: publicKey,
		Blockchain: BlockchainConfiguration{
			EthNodeURL:           ethNodeUrl,
			ContractAddress:      contractAddress,
			BlockchainPrivateKey: blockchainPrivateKey,
		},
	}
}

func stringFromEnv(key string, defaultValue string) string {
	var result, found = os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return result
}

func stringOrPanic(key string) string {
	var result, found = os.LookupEnv(key)
	if !found {
		panic(errors.New("configuration value not set for key: " + key))
	}
	return result
}

func intOrPanic(key string) int {
	var result, found = os.LookupEnv(key)

	if !found {
		panic(errors.New("configuration value not set for key: " + key))
	}

	intResult, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		panic(errors.New("configuration value for key: " + key + " is not a int"))
	}

	return int(intResult)
}
