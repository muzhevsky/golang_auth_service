package errs

import (
	"errors"
)

var UnexpectedError = errors.New("unexpected error")
var SomeErrorToDo = errors.New("not implemented yet")
var ValidationError = errors.New("validation error")
var DataBindError = errors.New("wrong data format")
var UnauthenticatedError = errors.New("authentication is necessary for this action")
var UserDataNotFoundError = errors.New("user data should be sent first")
var UserHasAlreadyPassedTestError = errors.New("user has already passed the test")
var UserDoesntHaveAnAvatarYetError = errors.New("user doesn't have an avatar yet")
