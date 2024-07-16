package tests

import (
	"smartri_app/internal/entities/user_data"
	"testing"
)

func TestXP(t *testing.T) {
	testCases := []struct {
		xp    user_data.XP
		valid bool
	}{
		{user_data.XP(0), true},
		{user_data.XP(14), true},
		{user_data.XP(24), true},
		{user_data.XP(44), true},
		{user_data.XP(94), true},
		{user_data.XP(1599), true},
		{user_data.XP(1600), true},
		{user_data.XP(1601), false},
		{user_data.XP(-1), false},
		{user_data.XP(-100), false},
		{user_data.XP(-1000 - 7), false},
	}

	for i := range testCases {
		err := testCases[i].xp.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].xp)
			return
		}
	}
}
