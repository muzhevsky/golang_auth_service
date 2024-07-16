package tests

import (
	"smartri_app/internal/entities/user_data_entities"
	"testing"
)

func TestXP(t *testing.T) {
	testCases := []struct {
		xp    user_data_entities.XP
		valid bool
	}{
		{user_data_entities.XP(0), true},
		{user_data_entities.XP(14), true},
		{user_data_entities.XP(24), true},
		{user_data_entities.XP(44), true},
		{user_data_entities.XP(94), true},
		{user_data_entities.XP(1599), true},
		{user_data_entities.XP(1600), true},
		{user_data_entities.XP(1601), false},
		{user_data_entities.XP(-1), false},
		{user_data_entities.XP(-100), false},
		{user_data_entities.XP(-1000 - 7), false},
	}

	for i := range testCases {
		err := testCases[i].xp.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].xp)
			return
		}
	}
}
