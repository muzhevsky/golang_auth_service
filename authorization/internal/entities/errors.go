package entities

import "errors"

var ValidationError = errors.New("validation error")

// verification
var ExpiredCode = errors.New("the verification code is already expired")
var UserIsNotVerified = errors.New("the user is not verified")
