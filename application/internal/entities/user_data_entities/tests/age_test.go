package tests

import (
	"smartri_app/internal/entities/user_data_entities"
	"testing"
)

func TestAge(t *testing.T) {
	testCases := []struct {
		age   user_data_entities.Age
		valid bool
	}{
		{user_data_entities.Age(4), true},
		{user_data_entities.Age(14), true},
		{user_data_entities.Age(24), true},
		{user_data_entities.Age(44), true},
		{user_data_entities.Age(94), true},
		{user_data_entities.Age(99), true},
		{user_data_entities.Age(1), false},
		{user_data_entities.Age(0), false},
		{user_data_entities.Age(-1), false},
		{user_data_entities.Age(100), false},
		{user_data_entities.Age(1000 - 7), false},
	}

	for i := range testCases {
		err := testCases[i].age.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].age)
			return
		}
	}
}
