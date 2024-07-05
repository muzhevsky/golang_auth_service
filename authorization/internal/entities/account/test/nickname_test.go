package test

import (
	"authorization/internal/entities/account"
	"testing"
)

type nicknameTestEntity struct {
	nickname account.Nickname
	valid    bool
}

func TestNickname(t *testing.T) {
	testCases := []nicknameTestEntity{
		{account.Nickname("this_nickname_is_toooooo_long"), false},
		{account.Nickname("a"), false},
		{account.Nickname("ПлохойНикнейм+"), false},
		{account.Nickname("HereIsQuestion?"), false},
		{account.Nickname("ХорошийНикнейм"), true},
		{account.Nickname("goodN1ckname"), true},
		{account.Nickname("1_is_good_nick"), true},
	}

	for i := range testCases {
		err := testCases[i].nickname.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].nickname)
			return
		}
	}
}
