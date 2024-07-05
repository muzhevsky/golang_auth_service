package test

import (
	"authorization/internal/entities/account"
	"testing"
)

type testPasswordEntity struct {
	password account.Password
	valid    bool
}

func TestPassword(t *testing.T) {
	testCases := []testPasswordEntity{
		{account.Password("thisPasswordIsToooooooLong"), false},
		{account.Password("short"), false},
		{account.Password("FinePassword123"), true},
		{account.Password("G00Dpa$$wOrD_321"), true},
		{account.Password("HereIs1QuestionMark?"), true},
		{account.Password("1_is_GoodDigit"), true},
	}

	for i := range testCases {
		err := testCases[i].password.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].password)
			return
		}
	}
}
