package requests

type UserDataRequest struct {
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type UserDataResponse struct {
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	XP     int    `json:"xp"`
}
