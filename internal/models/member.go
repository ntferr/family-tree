package models

type Member struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}

type Data struct {
	Name   string   `json:"name"`
	Member []Member `json:"member"`
}
