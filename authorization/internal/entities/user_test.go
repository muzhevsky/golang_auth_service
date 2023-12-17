package entities

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUser_isLetter(t *testing.T) {
	testTable := []struct {
		str      string
		expected bool
	}{
		{
			str:      "1sdf",
			expected: false,
		},
		{
			str:      "sdf",
			expected: true,
		},
		{
			str:      "-d",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := isLetter(rune(testCase.str[0]))
		t.Logf("Calling isLetter (%s), result %t\n", testCase.str, result)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %t, got %t", testCase.expected, result)
		}
	}
}

func TestUser_ValidateLogin(t *testing.T) {
	testTable := []struct {
		user     *User
		expected bool
	}{
		{
			user:     &User{Login: "dsf1"},
			expected: true,
		},
		{
			user:     &User{Login: "sd"},
			expected: false,
		},
		{
			user:     &User{Login: "2sdfd"},
			expected: false,
		},
		{
			user:     &User{Login: "sdf d"},
			expected: false,
		},
	}
	convertToString := func(expected bool) string {
		if expected {
			return "valid"
		}
		return "invalid"
	}

	for _, testCase := range testTable {
		err := testCase.user.ValidateLogin()
		t.Logf("Calling user.ValidateLogin(), result %s", err)

		if err == nil != testCase.expected {
			t.Errorf("Incorrect result. Login %s, must be %s", testCase.user.Login, convertToString(testCase.expected))
		}
	}
}

func TestUser_ValidateEmail(t *testing.T) {
	testTable := []struct {
		user     *User
		expected bool
	}{
		{
			user:     &User{EMail: "dima-piminov@yandex.ru"},
			expected: true,
		},
		{
			user:     &User{EMail: "dima-piminov@ya.ru"},
			expected: true,
		},
		{
			user:     &User{EMail: "2sdfd"},
			expected: false,
		},
		{
			user:     &User{EMail: "sdf d"},
			expected: false,
		},
		{
			user:     &User{EMail: "sdf d@sdf.com"},
			expected: false,
		},
	}
	convertToString := func(expected bool) string {
		if expected {
			return "valid"
		}
		return "invalid"
	}

	for _, testCase := range testTable {
		err := testCase.user.ValidateEmail()
		t.Logf("Calling user.ValidateEmail(), result %s", err)

		if err == nil != testCase.expected {
			t.Errorf("Incorrect result. e-mail %s, must be %s", testCase.user.EMail, convertToString(testCase.expected))
		}
	}
}

func TestUser_ValidatePassword(t *testing.T) {
	testTable := []struct {
		user     *User
		expected bool
	}{
		{
			user:     &User{Password: "gs123dfj"},
			expected: true,
		},
		{
			user:     &User{Password: "123231321"},
			expected: false,
		},
		{
			user:     &User{Password: "sfdfdsfsdfsd"},
			expected: false,
		},
		{
			user:     &User{Password: "1dfdf12"},
			expected: false,
		},
		{
			user:     &User{Password: "sdffsd2214214dsfsfds1"},
			expected: false,
		},
		{
			user:     &User{Password: "sdffsd2dfg@"},
			expected: false,
		},
	}
	convertToString := func(expected bool) string {
		if expected {
			return "valid"
		}
		return "invalid"
	}

	for _, testCase := range testTable {
		err := testCase.user.ValidatePassword()
		t.Logf("Calling user.ValidatePassword(), result %s", err)

		if err == nil != testCase.expected {
			t.Errorf("Incorrect result. password %s, must be %s", testCase.user.Password, convertToString(testCase.expected))
		}
	}
}
func TestUser_GenerateHashPassword(t *testing.T) {
	user := &User{Password: "sdfsfd"}
	err := user.GenerateHashPassword()
	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(user.Password))
	if err != nil {
		t.Errorf("CompareHashPassword failed!")
	}

	user = &User{Password: "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84chars"}
	err = user.GenerateHashPassword()
	if err == nil {
		t.Errorf("CompareHashPassword failed!")
	}
}
