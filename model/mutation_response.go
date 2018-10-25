package model

type MutationResponse struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}
