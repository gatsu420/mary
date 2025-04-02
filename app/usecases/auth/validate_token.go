package auth

import (
	"github.com/gatsu420/mary/common/errors"
	"github.com/golang-jwt/jwt/v4"
)

func (u *usecaseImpl) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, errors.New(errors.AuthError, "expected signing method HS256")
		}

		return []byte(u.secret), nil
	})
	if err != nil {
		return "", errors.New(errors.AuthError, "invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", errors.New(errors.AuthError, "failed to extract userID from claims")
		}

		return userID, nil
	}

	return "", errors.New(errors.AuthError, "invalid token")
}
