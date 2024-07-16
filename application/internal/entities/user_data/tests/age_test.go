package tests

import (
	"smartri_app/internal/entities/user_data"
	"testing"
)

func TestAge(t *testing.T) {
	testCases := []struct {
		age   user_data.Age
		valid bool
	}{
		{user_data.Age(4), true},
		{user_data.Age(14), true},
		{user_data.Age(24), true},
		{user_data.Age(44), true},
		{user_data.Age(94), true},
		{user_data.Age(99), true},
		{user_data.Age(1), false},
		{user_data.Age(0), false},
		{user_data.Age(-1), false},
		{user_data.Age(100), false},
		{user_data.Age(1000 - 7), false},
	}

	for i := range testCases {
		err := testCases[i].age.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].age)
			return
		}
	}
}
