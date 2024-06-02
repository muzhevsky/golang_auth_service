package requests

type VerificationRequest struct {
	UserId int    `json:"userId" binding:"required"`
	Code   string `json:"code" binding:"required"`
}
