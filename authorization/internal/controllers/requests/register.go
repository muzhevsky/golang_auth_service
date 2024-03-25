package requests

type CreateUserRequest struct {
	Login    string `json:"login" binding:"required" example:"TopPlayer123"`
	Password string `json:"password" binding:"required" example:"123superPassword"`
	EMail    string `json:"e-mail" binding:"required" example:"andrew123@qwerty.kom"`
	Nickname string `json:"nickname" binding:"required" example:"Looser1123"`
}
type CreateUserResponse struct {
	Id int `json:"id" example:"2"`
}
