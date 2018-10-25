package model

type Webhook struct {
	HashString string `json:"hashString"`
	GraphId    string `json:"graphId"`
	Type       string `json:"type"`
}
