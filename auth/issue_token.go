package auth

import (
	"time"

	"github.com/gatsu420/mary/common/errors"
	"github.com/golang-jwt/jwt/v4"
)

func (a *authImpl) IssueToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	signedToken, err := token.SignedString([]byte(a.secret))
	if err != nil {
		// TODO: pass err to new error in the same fashion as this:
		// return "", fmt.Errorf("unable to sign JWT: %w", err)
		return "", errors.New(errors.AuthError, "unable to sign JWT")
	}
	return signedToken, nil
}
