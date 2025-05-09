package auth

import (
	"context"
	"fmt"
	"hash/fnv"

	"github.com/gatsu420/mary/common/errors"
)

func (a *authImpl) CreateMembershipRegistry() ([]int, error) {
	// Bit slice should not be created everytime app starts, maybe do this instead:
	//	1.	generate slice using one-time worker and send to cache
	// 	2.	add bits when new user register

	users, err := a.query.ListUsers(context.Background())
	if err != nil {
		return nil, errors.New(errors.InternalServerError, "DB failed to list users")
	}

	// Length of slice and number of hash are determined by ballpark for now
	sliceLen := 100
	bits := make([]int, sliceLen)
	salts := []string{a.config.MembershipSalt1, a.config.MembershipSalt2, a.config.MembershipSalt3}

	hash := fnv.New64()
	var bitIdx int
	for _, u := range users {
		for _, s := range salts {
			hash.Write([]byte(fmt.Sprintf("%v %v %v", s, u, s)))
			bitIdx = abs(int(hash.Sum64()) % sliceLen)
			bits[bitIdx] = 1
			hash.Reset()
		}
	}
	return bits, nil
}

func (a *authImpl) CheckMembership(registry []int, username string) error {
	salts := []string{a.config.MembershipSalt1, a.config.MembershipSalt2, a.config.MembershipSalt3}
	hash := fnv.New64()
	var bitIdx int
	for _, s := range salts {
		hash.Write([]byte(fmt.Sprintf("%v %v %v", s, username, s)))
		bitIdx = abs(int(hash.Sum64()) % len(registry))
		if registry[bitIdx] == 0 {
			return errors.New(errors.AuthError, "user is not found")
		}
		hash.Reset()
	}
	return nil
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
