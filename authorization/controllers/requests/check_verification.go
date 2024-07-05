package requests

type CheckVerificationResponse struct {
	AccountId  int  `json:"accountId"`
	IsVerified bool `json:"isVerified"`
}
