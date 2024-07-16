package tests

import (
	"smartri_app/internal/entities/user_data"
	"testing"
)

func TestNickname(t *testing.T) {
	testCases := []struct {
		nickname user_data.Nickname
		valid    bool
	}{
		{user_data.Nickname("this_nickname_is_toooooo_long"), false},
		{user_data.Nickname("a"), false},
		{user_data.Nickname("ПлохойНикнейм+"), false},
		{user_data.Nickname("HereIsQuestion?"), false},
		{user_data.Nickname("ХорошийНикнейм"), true},
		{user_data.Nickname("goodN1ckname"), true},
		{user_data.Nickname("1_is_good_nick"), true},
	}

	for i := range testCases {
		err := testCases[i].nickname.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].nickname)
			return
		}
	}
}
