package errors

import "errors"

// validation
var LoginLengthError = errors.New("incorrect login length")
var LoginContentError = errors.New("incorrect characters in login")

// signUp
var EmailIsNotUnique = errors.New("user with specified email already exists")
var LoginIsNotUnique = errors.New("user with specified login already exists")
