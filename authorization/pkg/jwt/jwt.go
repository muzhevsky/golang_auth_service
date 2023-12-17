package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	signingString string
}

func New(signingString string) *JWT {
	return &JWT{
		signingString: signingString,
	}
}

func (j *JWT) NewToken(claim map[string]interface{}) (string, error) {
	var mapClaims jwt.MapClaims
	mapClaims = claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString([]byte(j.signingString))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
