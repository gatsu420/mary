package config

import "github.com/spf13/viper"

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresURL      string

	GRPCServerPort int

	JWTSecret string
}

func LoadConfig(filePath string) (*Config, error) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()

	return &Config{
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresPort:     viper.GetInt("POSTGRES_PORT"),
		PostgresDB:       viper.GetString("POSTGRES_DB"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresURL:      viper.GetString("POSTGRES_URL"),

		GRPCServerPort: viper.GetInt("GRPC_SERVER_PORT"),

		JWTSecret: viper.GetString("JWT_SECRET"),
	}, nil
}
