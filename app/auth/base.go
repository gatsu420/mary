package auth

import (
	"github.com/gatsu420/mary/common/config"
)

type Auth interface {
	IssueToken(username string) (string, error)
	ValidateToken(signedToken string) (string, error)
	CreateMembershipRegistry(users []string) []string
	CheckMembership(registry []string, username string) error
}

type authImpl struct {
	config *config.Config
}

var _ Auth = &authImpl{}

func NewAuth(config *config.Config) Auth {
	return &authImpl{
		config: config,
	}
}
