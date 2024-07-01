package errs

import (
	"errors"
)

var SomeErrorToDo = errors.New("not implemented yet")
var DataBindError = errors.New("wrong data format")
var UnauthenticatedError = errors.New("authenticated is necessary for this action")
var UserDataNotFoundError = errors.New("user data should be sent first")
var UserHasAlreadyPassedTestError = errors.New("user has already passed the test")
