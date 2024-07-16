package requests

type CloseSessionsRequest struct {
	Ids [16]int `json:"ids"`
}
