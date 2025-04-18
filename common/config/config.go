package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN    string
	GRPCServerPort string
	JWTSecret      string
}

func New(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return nil, err
	}

	return &Config{
		PostgresDSN:    os.Getenv("MARY_POSTGRES_DSN"),
		GRPCServerPort: os.Getenv("MARY_GRPC_SERVER_PORT"),
		JWTSecret:      os.Getenv("MARY_JWT_SECRET"),
	}, nil
}
