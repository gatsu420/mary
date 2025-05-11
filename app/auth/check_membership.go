package auth

import (
	"fmt"
	"hash/fnv"
	"math"

	"github.com/gatsu420/mary/common/errors"
)

func (a *authImpl) CheckMembership(registry []string, username string) error {
	salts := []string{a.config.MembershipSalt1, a.config.MembershipSalt2, a.config.MembershipSalt3}
	hash := fnv.New64()
	var bitIdx int
	var hashedVal uint64
	for _, s := range salts {
		hash.Write([]byte(fmt.Sprintf("%v %v %v", s, username, s)))
		hashedVal = hash.Sum64()
		bitIdx = int(math.Abs(float64(hashedVal))) % len(registry)
		if registry[bitIdx] == "0" {
			return errors.New(errors.AuthError, "user is not found")
		}
		hash.Reset()
	}
	return nil
}
