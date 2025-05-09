package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN     string
	CacheAddress    string
	GRPCServerPort  string
	JWTSecret       string
	MembershipSalt1 string
	MembershipSalt2 string
	MembershipSalt3 string
}

func New(filePath string) (*Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return nil, err
	}

	return &Config{
		PostgresDSN:     os.Getenv("MARY_POSTGRES_DSN"),
		CacheAddress:    os.Getenv("MARY_CACHE_ADDRESS"),
		GRPCServerPort:  os.Getenv("MARY_GRPC_SERVER_PORT"),
		JWTSecret:       os.Getenv("MARY_JWT_SECRET"),
		MembershipSalt1: os.Getenv("MARY_MEMBERSHIP_REGISTRY_SALT_1"),
		MembershipSalt2: os.Getenv("MARY_MEMBERSHIP_REGISTRY_SALT_2"),
		MembershipSalt3: os.Getenv("MARY_MEMBERSHIP_REGISTRY_SALT_3"),
	}, nil
}
