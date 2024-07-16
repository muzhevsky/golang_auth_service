package test

import (
	"authorization/internal/entities/account"
	"testing"
)

func TestEmail(t *testing.T) {
	testCases := []struct {
		email account.Email
		valid bool
	}{
		{account.Email("toolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolong@mail.ru"), false},
		{account.Email("a@a.a"), false},
		{account.Email("okmail@mail.ru"), true},
		{account.Email("Go0DmA1l@mail.ru"), true},
		{account.Email("smart+email@mail.ru"), true},
		{account.Email("HereIsAQuestionMark@mail.ru"), true},
		{account.Email("ярусскийДаТакТожеМожно@мда.рф"), true},
		{account.Email("its_not_an_email"), false},
		{account.Email("its_not_an_email_too@mail"), false},
		{account.Email("its_not_an_email_too@mail."), false},
		{account.Email("its_not_an_email_too@aaaaaaaaaaaaaaaaaaaaaaaaamail.ru"), false},
		{account.Email("its_not_an_email_too%№!!?*()(!?@aaaaaaaaaaaaaaaaaaaaaaaaamail."), false},
	}

	for i := range testCases {
		err := testCases[i].email.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].email)
			return
		}
	}
}
