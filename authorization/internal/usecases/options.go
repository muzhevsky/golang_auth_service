package usecases

import "time"

type sessionUseCaseOption func(useCase *SessionUseCase)

func NewAccessTokenDuration(duration time.Duration) sessionUseCaseOption {
	return func(useCase *SessionUseCase) {
		useCase.accessExpireDuration = duration
	}
}

func NewRefreshTokenDuration(duration time.Duration) sessionUseCaseOption {
	return func(useCase *SessionUseCase) {
		useCase.refreshExpireDuration = duration
	}
}
