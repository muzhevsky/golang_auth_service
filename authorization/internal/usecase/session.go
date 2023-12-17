package usecase

import "authorization/internal/entities"

type SessionUseCase struct {
	tokenGenerator ITokenGenerator
	sessionRepo    ISessionRepo
}

func NewSessionUseCase(generator ITokenGenerator, repo ISessionRepo) *SessionUseCase {
	return &SessionUseCase{generator, repo}
}

func (s SessionUseCase) CreateAccessToken(claims map[string]interface{}) (string, error) {
	return s.tokenGenerator.GenerateToken(claims)
}

func (s SessionUseCase) VerifyAccessToken(token string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionUseCase) CreateTokens(claims map[string]interface{}) (*entities.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionUseCase) UpdateAccessToken(accessToken, refreshToken string) (*entities.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionUseCase) AuthenticateUser(login, password string) (*entities.Session, error) {
	//TODO implement me
	panic("implement me")
}
