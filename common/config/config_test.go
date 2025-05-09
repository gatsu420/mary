package config_test

import (
	"testing"

	"github.com/gatsu420/mary/common/config"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	testCases := []struct {
		testName       string
		filePath       string
		expectedConfig *config.Config
	}{
		{
			testName:       "file is not found",
			filePath:       ".there.is.no.file.env",
			expectedConfig: nil,
		},
		{
			testName: "file is found",
			filePath: "../../.env.example",
			expectedConfig: &config.Config{
				PostgresDSN:     "postgres://user:pass@localhost:5432/testdb?sslmode=disable",
				CacheAddress:    "127.0.0.1:6379",
				GRPCServerPort:  "9090",
				JWTSecret:       "secret",
				MembershipSalt1: "salt1",
				MembershipSalt2: "salt2",
				MembershipSalt3: "salt3",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			cfg, err := config.New(tc.filePath)
			assert.Equal(t, tc.expectedConfig, cfg)
			if tc.expectedConfig == nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
