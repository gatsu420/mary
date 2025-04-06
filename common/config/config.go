package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresURL      string

	GRPCServerPort string

	JWTSecret string
}

func LoadConfig(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return nil, err
	}

	return &Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresURL:      os.Getenv("POSTGRES_URL"),

		GRPCServerPort: os.Getenv("GRPC_SERVER_PORT"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}, nil
}
