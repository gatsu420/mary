package auth_test

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *testSuite) Test_IssueToken() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "test",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	s.Run("secret is not []byte", func() {
		signedToken, err := token.SignedString("secret")
		s.Equal("", signedToken)
		s.NotNil(err)
	})

	s.Run("success", func() {
		signedToken, err := token.SignedString([]byte("secret"))
		s.NotEqual("", signedToken)
		s.Nil(err)
	})

	s.Run("ss", func() {
		signedToken, err := s.usecases.IssueToken("ss")
		s.NotEqual("", signedToken)
		s.Nil(err)
	})
}
