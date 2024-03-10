package entities

import (
	"errors"
)

var ValidationError = errors.New("validation error")

// auth
var UserNotFound = errors.New("there's no user with such email or login")
var WrongPassword = errors.New("wrong password")

var AccessTokenExpired = errors.New("access token is expired")
var RefreshTokenExpired = errors.New("refresh token is expired")
var NotAValidAccessToken = errors.New("invalid access token")
var NotAValidRefreshToken = errors.New("invalid refresh token")

// verification
var ExpiredCode = errors.New("the verification code is expired")
var WrongVerificationCode = errors.New("the verification code is wrong")
var UserIsNotVerified = errors.New("the user is not verified")
