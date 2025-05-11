package auth

import (
	"github.com/gatsu420/mary/app/repository"
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
	query  repository.Querier
}

var _ Auth = &authImpl{}

func NewAuth(config *config.Config, query repository.Querier) Auth {
	return &authImpl{
		config: config,
		query:  query,
	}
}
