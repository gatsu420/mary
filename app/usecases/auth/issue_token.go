package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (u *usecase) IssueToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	signedToken, _ := token.SignedString([]byte(u.secret))
	return signedToken
}
