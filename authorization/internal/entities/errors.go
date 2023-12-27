package entities

import "errors"

var ValidationError = errors.New("validation error")

// auth
var UserNotFound = errors.New("there's no user with such email or login")
var TokenExpired = errors.New("token is expired")
var WrongPassword = errors.New("wrong password")
var NotAValidToken = errors.New("invalid token")

// // verification
var ExpiredCode = errors.New("the verification code is already expired")
var UserIsNotVerified = errors.New("the user is not verified")
