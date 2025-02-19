package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Services interface {
	IssueToken(userID string) (string, error)
	ValidateToken(signedToken string) (string, error)
}

type service struct {
	secret string
}

func NewService(secret string) Services {
	return &service{secret}
}

func (s *service) IssueToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": time.Now().Unix(),
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
