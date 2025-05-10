package auth

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

func (a *authImpl) CreateMembershipRegistry(users []string) []string {
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

	strBits := make([]string, sliceLen)
	for i := range strBits {
		strBits[i] = strconv.Itoa(bits[i])
	}
	return strBits
}
