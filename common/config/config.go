package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL    string
	GRPCServerPort string
	JWTSecret      string
}

func LoadConfig(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return nil, err
	}

	return &Config{
		PostgresURL:    os.Getenv("POSTGRES_URL"),
		GRPCServerPort: os.Getenv("GRPC_SERVER_PORT"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
	}, nil
}
