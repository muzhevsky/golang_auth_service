package query_builders

import (
	"authorization/internal/entities/session_entities"
	sq "github.com/Masterminds/squirrel"
)

const (
	devicesTableName            = "devices"
	deviceIdFieldName           = "id"
	deviceAccountIdFieldName    = "account_id"
	deviceNameFieldName         = "name"
	deviceAccessTokenFieldName  = "session_access_token"
	deviceCreationDateFieldName = "session_creation_date"
)

func NewSelectDeviceByIdQuery(builder *sq.StatementBuilderType, id int) (string, []any, error) {
	return selectStarDevice(builder).
		Where(sq.Eq{deviceIdFieldName: id}).
		ToSql()
}

func NewSelectDeviceByAccessTokenQuery(builder *sq.StatementBuilderType, accessToken string) (string, []any, error) {
	return selectStarDevice(builder).
		Where(sq.Eq{deviceAccessTokenFieldName: accessToken}).
		ToSql()
}

func NewSelectDeviceByAccountIdQuery(builder *sq.StatementBuilderType, accountId int) (string, []any, error) {
	return selectStarDevice(builder).
		Where(sq.Eq{deviceAccountIdFieldName: accountId}).
		ToSql()
}

func NewInsertDeviceQuery(builder *sq.StatementBuilderType, device *session_entities.Device) (string, []any, error) {
	return builder.
		Insert(devicesTableName).
		Columns(deviceAccountIdFieldName, deviceNameFieldName, deviceAccessTokenFieldName, deviceCreationDateFieldName).
		Values(device.AccountId, device.Name, device.SessionAccessToken, device.SessionCreationTime).
		ToSql()
}

func NewUpdateDeviceByAccessTokenQuery(builder *sq.StatementBuilderType, accessToken string, newDevice *session_entities.Device) (string, []any, error) {
	return builder.
		Update(devicesTableName).
		Set(deviceAccountIdFieldName, newDevice.AccountId).
		Set(deviceNameFieldName, newDevice.Name).
		Set(deviceAccessTokenFieldName, newDevice.SessionAccessToken).
		Set(deviceCreationDateFieldName, newDevice.SessionCreationTime).
		Where(sq.Eq{deviceAccessTokenFieldName: accessToken}).
		ToSql()
}

func NewDeleteDeviceByIdQuery(builder *sq.StatementBuilderType, id int) (string, []any, error) {
	return builder.
		Delete(devicesTableName).
		Where(sq.Eq{deviceIdFieldName: id}).
		ToSql()
}

func selectStarDevice(builder *sq.StatementBuilderType) sq.SelectBuilder {
	return builder.Select(
		deviceIdFieldName,
		deviceAccountIdFieldName,
		deviceNameFieldName,
		deviceAccessTokenFieldName,
		deviceCreationDateFieldName).
		From(devicesTableName)
}
