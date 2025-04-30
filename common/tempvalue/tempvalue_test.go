package tempvalue_test

import (
	"testing"

	"github.com/gatsu420/mary/common/tempvalue"
	"github.com/stretchr/testify/assert"
)

func Test_SetAndGetUserID(t *testing.T) {
	t.Run("set and get userID successfully", func(t *testing.T) {
		fakeUserID := "fakeUserID"

		tempvalue.SetUserID(fakeUserID)
		userID := tempvalue.GetUserID()
		assert.Equal(t, fakeUserID, userID)
	})
}

func Test_SetAndGetCalledMethods(t *testing.T) {
	t.Run("set and get calledMethods successfully", func(t *testing.T) {
		fakeMethod := "fake method"
		tempvalue.SetCalledMethods(fakeMethod)
		dummyMethod := "dummy method"
		tempvalue.SetCalledMethods(dummyMethod)

		methods := tempvalue.GetCalledMethods()
		_, ok := methods[fakeMethod]
		assert.True(t, ok)
		_, ok = methods[dummyMethod]
		assert.True(t, ok)
		_, ok = methods["non existent method"]
		assert.False(t, ok)
	})
}

func Test_SetAndFlushCalledMethods(t *testing.T) {
	t.Run("set and flush calledMethods successfully", func(t *testing.T) {
		fakeMethod := "fake method"
		tempvalue.SetCalledMethods(fakeMethod)
		dummyMethod := "dummy method"
		tempvalue.SetCalledMethods(dummyMethod)

		tempvalue.FlushCalledMethods()
		assert.Equal(t, 0, len(tempvalue.GetCalledMethods()))
	})
}
