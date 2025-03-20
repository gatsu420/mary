package auth

import (
	"context"
	"time"

	"github.com/gatsu420/mary/common/errors"
	"github.com/golang-jwt/jwt/v4"
)

func (u *usecase) IssueToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	signedToken, err := token.SignedString([]byte(u.secret))
	if err != nil {
		// TODO: pass err to new error in the same fashion as this:
		// return "", fmt.Errorf("unable to sign JWT: %w", err)
		return "", errors.New(errors.AuthError, "unable to sign JWT")
	}
	return signedToken, nil
}

func (u *usecase) ValidateToken(signedToken string) (string, error) {
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

func (u *usecase) CheckUserIsExisting(ctx context.Context, username string) error {
	isExisting, err := u.query.CheckUserIsExisting(ctx, username)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to check if user is existing")
	}
	if !isExisting {
		return errors.New(errors.NotFoundError, "user not found")
	}
	return nil
}
