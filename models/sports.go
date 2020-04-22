package models

//Sport representa os esportes na chamada  api
type Sport struct {
	Key          string `json:"key"`
	Active       bool   `json:"active" `
	Group        string `json:"group"`
	Details      string `json:"details"`
	Title        string `json:"title"`
	HasOutrights bool   `json:"has_outrights"`
}

//Response é o objeto de recebimento dos dados
type Response struct {
	Success string `json:"success"`
	Data    []Sport
}
