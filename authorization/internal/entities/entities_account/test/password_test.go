package test

import (
	"authorization/internal/entities/entities_account"
	"testing"
)

type testPasswordEntity struct {
	password entities_account.Password
	valid    bool
}

func TestPassword(t *testing.T) {
	testCases := []testPasswordEntity{
		{entities_account.Password("thisPasswordIsToooooooLong"), false},
		{entities_account.Password("short"), false},
		{entities_account.Password("FinePassword123"), true},
		{entities_account.Password("G00Dpa$$wOrD_321"), true},
		{entities_account.Password("HereIs1QuestionMark?"), true},
		{entities_account.Password("1_is_GoodDigit"), true},
	}

	for i := range testCases {
		err := testCases[i].password.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].password)
			return
		}
	}
}
