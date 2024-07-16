package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"context"
	"time"
)

type sessionRepository struct {
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand
	insertSessionCommand              datasources.IInsertSessionCommand
	updateSessionByAccessTokenCommand datasources.IUpdateSessionByAccessTokenCommand
	selectDeviceByAccessTokenCommand  datasources.ISelectDeviceByAccessTokenCommand
	insertDeviceCommand               datasources.IInsertDeviceCommand
	updateDeviceByAccessTokenCommand  datasources.IUpdateDeviceByAccessTokenCommand
}

func NewSessionRepository(
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand,
	insertSessionCommand datasources.IInsertSessionCommand,
	updateSessionByAccessTokenCommand datasources.IUpdateSessionByAccessTokenCommand,
	selectDeviceByAccessTokenCommand datasources.ISelectDeviceByAccessTokenCommand,
	insertDeviceCommand datasources.IInsertDeviceCommand,
	updateDeviceByAccessTokenCommand datasources.IUpdateDeviceByAccessTokenCommand) internal.ISessionRepository {

	return &sessionRepository{
		selectSessionByAccessTokenCommand: selectSessionByAccessTokenCommand,
		insertSessionCommand:              insertSessionCommand,
		updateSessionByAccessTokenCommand: updateSessionByAccessTokenCommand,
		selectDeviceByAccessTokenCommand:  selectDeviceByAccessTokenCommand,
		insertDeviceCommand:               insertDeviceCommand,
		updateDeviceByAccessTokenCommand:  updateDeviceByAccessTokenCommand}
}

func (s *sessionRepository) CreateWithDevice(context context.Context, deviceName string, se *session.Session) error {
	device := session.Device{
		AccountId:           se.AccountId,
		Name:                deviceName,
		SessionAccessToken:  se.AccessToken,
		SessionCreationTime: time.Now(),
	}
	err := s.insertDeviceCommand.Execute(context, &device)
	if err != nil {
		return err
	}
	return s.insertSessionCommand.Execute(context, se)
}

func (s *sessionRepository) FindByAccessToken(context context.Context, token string) (*session.Session, error) {
	return s.selectSessionByAccessTokenCommand.Execute(context, token)
}

func (s *sessionRepository) UpdateByAccessToken(context context.Context, token string, newSession *session.Session) (*session.Session, error) {
	cpy := session.Session{
		AccountId:    newSession.AccountId,
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiresAt:    newSession.ExpiresAt,
	}

	device, err := s.selectDeviceByAccessTokenCommand.Execute(context, token)
	if err != nil {
		return nil, err
	}

	device.SessionAccessToken = cpy.AccessToken
	err = s.updateDeviceByAccessTokenCommand.Execute(context, token, device)
	if err != nil {
		return nil, err
	}

	err = s.updateSessionByAccessTokenCommand.Execute(context, token, &cpy)
	if err != nil {
		return nil, err
	}

	return newSession, err
}
