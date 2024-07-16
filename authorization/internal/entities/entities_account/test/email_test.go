package test

import (
	"authorization/internal/entities/entities_account"
	"testing"
)

func TestEmail(t *testing.T) {
	testCases := []struct {
		email entities_account.Email
		valid bool
	}{
		{entities_account.Email("toolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolong@mail.ru"), false},
		{entities_account.Email("a@a.a"), false},
		{entities_account.Email("okmail@mail.ru"), true},
		{entities_account.Email("Go0DmA1l@mail.ru"), true},
		{entities_account.Email("smart+email@mail.ru"), true},
		{entities_account.Email("HereIsAQuestionMark@mail.ru"), true},
		{entities_account.Email("ярусскийДаТакТожеМожно@мда.рф"), true},
		{entities_account.Email("its_not_an_email"), false},
		{entities_account.Email("its_not_an_email_too@mail"), false},
		{entities_account.Email("its_not_an_email_too@mail."), false},
		{entities_account.Email("its_not_an_email_too@aaaaaaaaaaaaaaaaaaaaaaaaamail.ru"), false},
		{entities_account.Email("its_not_an_email_too%№!!?*()(!?@aaaaaaaaaaaaaaaaaaaaaaaaamail."), false},
	}

	for i := range testCases {
		err := testCases[i].email.Validate()
		if testCases[i].valid != (err == nil) {
			t.Fatalf("Test case \"%v\" failed\n", testCases[i].email)
			return
		}
	}
}
