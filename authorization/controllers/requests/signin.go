package requests

type SignInRequest struct {
	Login      string `json:"login" example:"TopPlayer123"`
	Password   string `json:"password" example:"123superPassword"`
	DeviceName string `json:"device_name" example:"Google Pixel 8a"`
}

type SignInResponse struct {
	Id      int                     `json:"id" example:"2"`
	Session *RefreshSessionResponse `json:"session"`
}

func NewSignInResponse(id int, session *RefreshSessionResponse) *SignInResponse {
	return &SignInResponse{Id: id, Session: session}
}
