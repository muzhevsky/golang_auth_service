package infrastructure

import "authorization/pkg/jwt"

type JWTGenerator struct {
}

func (generator *JWTGenerator) GenerateToken(claims map[string]interface{}) (string, error) {
	return (&jwt.JWT{}).NewToken(claims)
}
