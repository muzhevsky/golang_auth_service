package tests

import (
	"smartri_app/internal/entities/user_data"
	"testing"
)

func TestGender(t *testing.T) {
	testCases := []struct {
		gender user_data.Gender
		valid  bool
	}{
		{user_data.Gender("f"), true},
		{user_data.Gender("m"), true},
		{user_data.Gender("n"), false},
		{user_data.Gender("test123"), false},
		{user_data.Gender(""), false},
		{user_data.Gender("1"), false},
	}

	for i := range testCases {
		err := testCases[i].gender.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].gender)
			return
		}
	}
}
