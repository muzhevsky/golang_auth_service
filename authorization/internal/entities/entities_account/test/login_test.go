package test

import (
	"authorization/internal/entities/entities_account"
	"testing"
)

type testLoginEntity struct {
	login entities_account.Login
	valid bool
}

func TestLogin(t *testing.T) {
	testCases := []testLoginEntity{
		{entities_account.Login("this_login_consists_of_31_chars"), false},
		{entities_account.Login("a"), false},
		{entities_account.Login("fineLogin123"), true},
		{entities_account.Login("goodLogin123"), true},
		{entities_account.Login("HereIsAQuestionMark?"), false},
		{entities_account.Login("1_is_forbidden"), false},
	}

	for i := range testCases {
		err := testCases[i].login.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].login)
			return
		}
	}
}
