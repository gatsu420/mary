package auth

import (
	"context"
	"time"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/golang-jwt/jwt/v4"
)

type Services interface {
	IssueToken(username string) (string, *errors.Error)
	ValidateToken(signedToken string) (string, *errors.Error)
	CheckUserIsExisting(ctx context.Context, username string) *errors.Error
}

type service struct {
	secret string
	query  *repository.Queries
}

func NewService(secret string, query *repository.Queries) Services {
	return &service{
		secret: secret,
		query:  query,
	}
}

func (s *service) IssueToken(username string) (string, *errors.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		// return "", fmt.Errorf("unable to sign JWT: %w", err)
		return "", errors.Auth("unable to sign JWT")
	}

	return signedToken, nil
}

func (s *service) ValidateToken(signedToken string) (string, *errors.Error) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			// return nil, fmt.Errorf("expected signing method: %v", t.Header["alg"])
			return nil, errors.Auth("expected signing method HS256")
		}

		return []byte(s.secret), nil
	})
	if err != nil {
		// return "", fmt.Errorf("invalid token")
		return "", errors.Auth("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			// return "", fmt.Errorf("failed to extract userID from claims")
			return "", errors.Auth("failed to extract userID from claims")
		}

		return userID, nil
	}

	// return "", fmt.Errorf("invalid token")
	return "", errors.Auth("invalid token")
}

func (s *service) CheckUserIsExisting(ctx context.Context, username string) *errors.Error {
	isExisting, err := s.query.CheckUserIsExisting(ctx, username)
	if err != nil {
		return errors.InternalServer("DB failed to check if user is existing")
	}
	if !isExisting {
		return errors.NotFound("user is not found")
	}

	return nil
}
