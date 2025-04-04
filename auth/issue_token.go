package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (a *authImpl) IssueToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	// SignedString will still return token eventhouggh secret is empty string.
	signedToken, _ := token.SignedString([]byte(a.secret))
	return signedToken, nil
}
