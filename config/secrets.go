package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

const (
	ENV_FILE = ".env"
)

type Secrets struct {
	Env              string `env:"ENV"`
	MongoDBHost      string `env:"MONGO_HOST"`
	MongoDBPassword  string `env:"MONGO_PASSWORD"`
	MongoDBUsername  string `env:"MONGO_USERNAME"`
	LocalMongoDBHost string `env:"LOCAL_MONGO_DB_HOST"`
}

func InitializeSecrets() *Secrets {
	logger = GetLogger("secrets")
	secrets := parseEnv()
	return secrets
}

func loadEnvFile() error {
	// try to load env file, if not exist return error else load is done.
	if _, err := os.Stat(ENV_FILE); os.IsNotExist(err) {
		return err
	}
	return godotenv.Load(ENV_FILE)
}

func parseEnv() *Secrets {
	if err := loadEnvFile(); err != nil {
		logger.Errorf("erro to load dotenv file, error: %v", err)
		return &Secrets{}
	}
	sct := &Secrets{}
	if err := env.Parse(sct); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return sct
}
