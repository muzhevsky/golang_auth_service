package test

import (
	"authorization/internal/entities/account"
	"testing"
)

type testLoginEntity struct {
	login account.Login
	valid bool
}

func TestLogin(t *testing.T) {
	testCases := []testLoginEntity{
		{account.Login("this_login_consists_of_31_chars"), false},
		{account.Login("a"), false},
		{account.Login("fineLogin123"), true},
		{account.Login("goodLogin123"), true},
		{account.Login("HereIsAQuestionMark?"), false},
		{account.Login("1_is_forbidden"), false},
	}

	for i := range testCases {
		err := testCases[i].login.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].login)
			return
		}
	}
}
