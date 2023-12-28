package usecase

//
//import (
//	"authorization/internal/entities"
//	"context"
//)
//
//type securityRoles struct {
//	userRepo         IUserRepo
//	rolesRepo        IRoleRepo
//	routeSecurityMap map[string][][]entities.Role
//}
//
//func NewSecurityUseCase(userRepo IUserRepo, rolesRepo IRoleRepo, routeSecurityMap map[string][][]string) ISecurity {
//
//	roleMap := make(map[string][][]entities.Role)
//	for key, value := range routeSecurityMap {
//		roles := make([][]entities.Role, len(value))
//		for innerKey, innerValue := range value {
//			roleSet := make([]entities.Role, len(value))
//			for i := 0; i < len(roleSet); i++ {
//				roleSet[0] = entities.Role()
//			}
//		}
//	}
//	return &securityRoles{userRepo: userRepo, rolesRepo: rolesRepo, routeSecurityMap: routeSecurityMap}
//}
//
//func (s *securityRoles) CheckAccess(ctx context.Context, route string, userId int) (bool, error) {
//	if s.routeSecurityMap[route] == nil {
//		return true, nil
//	}
//
//	roles, err := s.rolesRepo.GetRoles(ctx, userId)
//	if err != nil {
//		return false, err
//	}
//
//	rolesChecksum := 0
//	for role := range roles {
//		rolesChecksum += role
//	}
//
//	for roleSet := range s.routeSecurityMap {
//		roleSetCheckSum := 0
//		for role := range roleSet {
//			roleSetCheckSum += role
//		}
//		if rolesChecksum == roleSetCheckSum {
//			return true, nil
//		}
//	}
//	return false, nil
//}
