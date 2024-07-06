package requests

type SignInRequest struct {
	Login    string `json:"login" example:"TopPlayer123"`
	Password string `json:"password" example:"123superPassword"`
}

type SignInResponse struct {
	Id      int                    `json:"id" example:"2"`
	Session RefreshSessionResponse `json:"session"`
}
