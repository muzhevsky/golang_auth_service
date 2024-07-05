package requests

type VerificationRequest struct {
	Code string `json:"code" binding:"required" example:"a1b2c"`
}
