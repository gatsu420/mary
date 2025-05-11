package auth_test

func (s *testSuite) Test_CreateMembershipRegistry() {
	s.Run("create membership registry successfully", func() {
		registry := s.auth.CreateMembershipRegistry([]string{"testUserA", "testUserB"})

		// Bootstrap registry to deduce which index has value 1 or 0. As long as salts in
		// .env.example don't change, these assertions will work.
		s.Equal("1", registry[14])
		s.Equal("1", registry[30])
		s.Equal("0", registry[31])
		s.Equal("0", registry[32])
	})
}
