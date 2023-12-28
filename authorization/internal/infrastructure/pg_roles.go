package infrastructure

//
//import (
//	"authorization/internal/entities"
//	"authorization/internal/usecase"
//	"authorization/pkg/postgres"
//	"context"
//	sq "github.com/Masterminds/squirrel"
//)
//
//type rolesRepo struct {
//	pg *postgres.Postgres
//}
//
//func NewRoleRepo(pg *postgres.Postgres) usecase.IRoleRepo {
//	return &rolesRepo{pg}
//}
//
//func (r *rolesRepo) GetRoles(ctx context.Context, userId int) ([]entities.Role, error) {
//	sql, args, err := r.pg.Builder.Select("roleId").
//		From("users_and_roles").
//		Where(sq.Eq{"userId": userId}).
//		ToSql()
//
//	if err != nil {
//		return nil, err
//	}
//
//	rows, err := r.pg.Pool.Query(ctx, sql, args...)
//	if err != nil {
//		return nil, err
//	}
//
//	roles := make([]entities.Role, 4)
//	for rows.Next() {
//		roleId := 0
//		err = rows.Scan(&roleId)
//		if err != nil {
//			return nil, err
//		}
//
//		roleName, err := r.getRoleName(ctx, roleId)
//		if err != nil {
//			return nil, err
//		}
//
//		role, err := entities.StringToRole(roleName)
//		if err != nil {
//			return nil, err
//		}
//
//		roles = append(roles, role)
//	}
//
//	return roles, nil
//}
//
//func (r *rolesRepo) getRoleName(ctx context.Context, roleId int) (string, error) {
//	sql, args, err := r.pg.Builder.Select("name").
//		From("users_and_roles").
//		Where(sq.Eq{"userId": roleId}).
//		ToSql()
//
//	if err != nil {
//		return "", err
//	}
//
//	result := ""
//	err = r.pg.Pool.QueryRow(ctx, sql, args).Scan(&result)
//	if err != nil {
//		return "", err
//	}
//
//	return result, nil
//}
//
//func (r *rolesRepo) SetRoles(ctx context.Context, userId int, roles []entities.Role) error {
//	//TODO implement me
//	panic("implement me")
//}
