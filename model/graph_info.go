package model

type GraphInfo struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Unit           string `json:"unit"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	PurgeCacheURLs string `json:"purgeCacheURLs"`
}
