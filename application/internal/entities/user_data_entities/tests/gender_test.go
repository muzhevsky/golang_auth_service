package tests

import (
	"smartri_app/internal/entities/user_data_entities"
	"testing"
)

func TestGender(t *testing.T) {
	testCases := []struct {
		gender user_data_entities.Gender
		valid  bool
	}{
		{user_data_entities.Gender("f"), true},
		{user_data_entities.Gender("m"), true},
		{user_data_entities.Gender("n"), false},
		{user_data_entities.Gender("test123"), false},
		{user_data_entities.Gender(""), false},
		{user_data_entities.Gender("1"), false},
	}

	for i := range testCases {
		err := testCases[i].gender.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].gender)
			return
		}
	}
}
