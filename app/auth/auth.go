package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/db/repository"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Services interface {
	IssueToken(username string) (string, error)
	ValidateToken(signedToken string) (string, error)
	CheckUserIsExisting(ctx context.Context, username string) error
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

func (s *service) IssueToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", fmt.Errorf("unable to sign JWT: %w", err)
	}

	return signedToken, nil
}

func (s *service) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("expected signing method: %v", t.Header["alg"])
		}

		return []byte(s.secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("failed to extract userID from claims")
		}

		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}

func (s *service) CheckUserIsExisting(ctx context.Context, username string) error {
	isExisting, err := s.query.CheckUserIsExisting(ctx, username)
	if err != nil {
		return err
	}
	if !isExisting {
		return status.Error(codes.NotFound, "user not found")
	}

	return nil
}
