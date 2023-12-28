package usecase

//
//import (
//	"context"
//)
//
//type securityVerification struct {
//	userRepo                   IUserRepo
//	verificationRequiredRoutes []string
//}
//
//func NewSecurity(userRepo IUserRepo, shouldBeVerified []string) ISecurity {
//	return &securityVerification{userRepo: userRepo, verificationRequiredRoutes: shouldBeVerified}
//}
//
//func (s *securityVerification) CheckAccess(ctx context.Context, route string, userId int) (bool, error) {
//	for _, storedRoute := range s.verificationRequiredRoutes {
//		if route == storedRoute {
//			result, err := s.check(ctx, userId)
//			if err != nil {
//				return false, err
//			}
//			return result, nil
//		}
//	}
//	return true, nil
//}
//
//func (s *securityVerification) check(ctx context.Context, userId int) (bool, error) {
//	user, err := s.userRepo.FindById(ctx, userId)
//	if err != nil {
//		return false, err
//	}
//
//	return user.IsVerified, nil
//}
