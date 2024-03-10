package infrastructure

import (
	"authorization/pkg/jwt"
)

type jwtProvider struct {
	jwt *jwt.JWT
}

func NewJwtProvider(jwt *jwt.JWT) *jwtProvider {
	return &jwtProvider{jwt: jwt}
}

func (provider *jwtProvider) GenerateToken(claims map[string]interface{}) (string, error) {
	return provider.jwt.NewToken(claims)
}

func (provider *jwtProvider) ParseToken(token string) (map[string]interface{}, error) {
	return provider.jwt.ParseToken(token)
}
