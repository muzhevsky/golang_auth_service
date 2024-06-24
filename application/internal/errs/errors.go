package errs

import (
	"errors"
)

var SomeErrorToDo = errors.New("not implemented yet")
var DataBindError = errors.New("wrong data format")
var UnauthenticatedError = errors.New("authenticated is necessary for this action")
