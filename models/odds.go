package models

import (
	"milhonarios/utils"
)

//Odds pertence a sites
type Odds struct {
	H2H    []float32 `json:"h2h"`
	H2HLay []float32 `json:"h2h_lay"`
}

//Sites é o  objeto que recebe da api sites
type Sites struct {
	SiteKey    string     `json:"site_key"`
	SiteNice   string     `json:"site_nice"`
	LastUpdate utils.Time `json:"last_update"`
	Odds       Odds       `json:"odds"`
}

//Odd é o objeto que recebe odds
type Odd struct {
	Sportkey     string     `json:"sport_key"`
	SportNice    string     `json:"sport_nice"`
	Teams        []string   `json:"teams"`
	CommenceTime utils.Time `json:"commence_time"`
	HomeTeam     string     `json:"home_team"`
	Sites        []Sites    `json:"sites"`
	SitesCount   int8       `json:"sites_count"`
}

//OddsResponse é o objeto de recebimento dos dados
type OddsResponse struct {
	Success bool  `json:"success"`
	Data    []Odd `json:"data"`
}

//Filter filtra algo
func (vs *OddsResponse) Filter(f func(criteria Odd) bool) []Odd {
	vsf := make([]Odd, 0)

	for _, v := range vs.Data {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
