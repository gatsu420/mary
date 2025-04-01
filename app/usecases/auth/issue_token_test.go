package auth_test

func (s *testSuite) Test_IssueToken() {
	s.Run("success", func() {
		signedToken := s.usecases.IssueToken("test")
		s.NotEqual("", signedToken)
	})
}
