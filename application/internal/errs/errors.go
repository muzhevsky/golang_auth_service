package errs

import (
	"errors"
)

var SomeErrorToDo = errors.New("not implemented yet")
var ValidationError = errors.New("validation error")
var DataBindError = errors.New("wrong data format")
var UnauthenticatedError = errors.New("authentication is necessary for this action")
var UserDataNotFoundError = errors.New("user data should be sent first")
var UserHasAlreadyPassedTestError = errors.New("user has already passed the test_entities")
var UserDoesntHaveAnAvatarYetError = errors.New("user doesn't have an avatar_entities yet")
