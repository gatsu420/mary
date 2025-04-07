package auth

import (
	"github.com/gatsu420/mary/common/errors"
	"github.com/golang-jwt/jwt/v4"
)

func (a *authImpl) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})
	if err != nil {
		return "", errors.New(errors.AuthError, "unable to parse or verify token")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	userID, _ := claims["sub"].(string)

	return userID, nil
}
