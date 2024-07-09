package query_builders

import (
	"authorization/internal/entities/session"
	sq "github.com/Masterminds/squirrel"
)

const (
	devicesTableName            = "devices"
	deviceIdFieldName           = "id"
	deviceIdentityFieldName     = "identity"
	deviceNameFieldName         = "name"
	deviceAccessTokenFieldName  = "session_access_token"
	deviceCreationDateFieldName = "session_creation_date"
)

func NewSelectDeviceByIdQuery(builder *sq.StatementBuilderType, id int) (string, []any, error) {
	return selectStarDevice(builder).
		Where(sq.Eq{deviceIdFieldName: id}).
		ToSql()
}

func NewSelectDeviceByIdentityQuery(builder *sq.StatementBuilderType, identity string) (string, []any, error) {
	return selectStarDevice(builder).
		Where(sq.Eq{deviceIdentityFieldName: identity}).
		ToSql()
}

func NewCreateDeviceQuery(builder *sq.StatementBuilderType, device *session.Device) (string, []any, error) {
	return builder.
		Insert(devicesTableName).
		Columns(deviceIdentityFieldName, deviceNameFieldName, deviceAccessTokenFieldName, deviceCreationDateFieldName).
		Values(device.Identity, device.Name, device.SessionAccessToken, device.SessionCreationTime).
		ToSql()
}

func NewUpdateDeviceByIdQuery(builder *sq.StatementBuilderType, id int, newDevice *session.Device) (string, []any, error) {
	return builder.
		Update(devicesTableName).
		Set(deviceIdentityFieldName, newDevice.Identity).
		Set(deviceNameFieldName, newDevice.Name).
		Set(deviceAccessTokenFieldName, newDevice.SessionAccessToken).
		Set(deviceCreationDateFieldName, newDevice.SessionCreationTime).
		Where(sq.Eq{deviceIdFieldName: id}).
		ToSql()
}

func selectStarDevice(builder *sq.StatementBuilderType) sq.SelectBuilder {
	return builder.Select(deviceIdFieldName, deviceIdentityFieldName, deviceNameFieldName, deviceAccessTokenFieldName, deviceCreationDateFieldName).
		From(devicesTableName)
}
