package models

//Sport representa os esportes na chamada  api
type sport struct {
	Key          string `json:"key"`
	Active       bool   `json:"active" `
	Group        string `json:"group"`
	Details      string `json:"details"`
	Title        string `json:"title"`
	HasOutrights bool   `json:"has_outrights"`
}

//SportsResponse é o objeto de recebimento dos dados
type SportsResponse struct {
	Success bool    `json:"success"`
	Data    []sport `json:"data"`
}
