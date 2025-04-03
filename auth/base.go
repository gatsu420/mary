package auth

type Auth interface {
	IssueToken(username string) (string, error)
	ValidateToken(signedToken string) (string, error)
}

type authImpl struct {
	secret string
}

func NewAuth(secret string) Auth {
	return &authImpl{
		secret: secret,
	}
}
