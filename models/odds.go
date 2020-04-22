package models

import "time"

//Odds pertence a sites
type odds struct {
	H2H    []float32 `json:"h2h"`
	H2HLay []float32 `json:"h2h_lay"`
}

//Sites é o  objeto que recebe da api sites
type sites struct {
	SiteKey    string    `json:"site_key"`
	SiteNice   string    `json:"site_nice"`
	LastUpdate time.Time `json:"last_update"`
	Odds       odds      `json:"odds"`
}

//Odd é o objeto que recebe odds
type odd struct {
	Sportkey     string    `json:"sport_key"`
	SportNice    string    `json:"sport_nice"`
	Teams        []string  `json:"teams"`
	CommenceTime time.Time `json:"commence_time"`
	Sites        []sites
}

//OddsResponse é o objeto de recebimento dos dados
type OddsResponse struct {
	Success string `json:"success"`
	Data    []odd
}
