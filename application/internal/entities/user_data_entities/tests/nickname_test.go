package tests

import (
	"smartri_app/internal/entities/user_data_entities"
	"testing"
)

func TestNickname(t *testing.T) {
	testCases := []struct {
		nickname user_data_entities.Nickname
		valid    bool
	}{
		{user_data_entities.Nickname("this_nickname_is_toooooo_long"), false},
		{user_data_entities.Nickname("a"), false},
		{user_data_entities.Nickname("ПлохойНикнейм+"), false},
		{user_data_entities.Nickname("HereIsQuestion?"), false},
		{user_data_entities.Nickname("ХорошийНикнейм"), true},
		{user_data_entities.Nickname("goodN1ckname"), true},
		{user_data_entities.Nickname("1_is_good_nick"), true},
	}

	for i := range testCases {
		err := testCases[i].nickname.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].nickname)
			return
		}
	}
}
