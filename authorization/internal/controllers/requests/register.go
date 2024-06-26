package requests

type CreateAccountRequest struct {
	Login    string `json:"login" binding:"required" example:"TopPlayer123"`
	Password string `json:"password" binding:"required" example:"123superPassword"`
	EMail    string `json:"e-mail" binding:"required" example:"andrew123@qwerty.kom"`
	Nickname string `json:"nickname" binding:"required" example:"SlimShady123"`
}
type CreateAccountResponse struct {
	Id      int                    `json:"id" example:"2"`
	Session RefreshSessionResponse `json:"session"`
}
