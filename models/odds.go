package models

import (
	"fmt"
	"milhonarios/utils"
)

//Odds pertence a sites
type Odds struct {
	H2H []float32 `json:"h2h"`
	//H2HLay []float32 `json:"h2h_lay"`
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

//Filter filtra
func (vs *OddsResponse) Filter(f func(criteria Odd) bool) []Odd {
	vsf := make([]Odd, 0)

	for _, v := range vs.Data {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

type Analisados struct {
	Data       int
	SitesIndex []int
}

//FilterSites filtra
func (vs *OddsResponse) FilterSites(f func(criteria Odd) bool) []Sites {
	sts := make([]Sites, 0)
	analisados := make([]Analisados, 0)

	for index, v := range vs.Data {
		var itemAnalisado Analisados
		if f(v) {
			itemAnalisado.Data = index
			for siteindex, x := range v.Sites {
				sts = append(sts, x)
				siteComparar := siteindex

				for indexOdd, oddAnterior := range x.Odds.H2H {
					s := fmt.Sprintf("%f >>", oddAnterior)
					//comparar com  o proximo Odd do proximo site
					siteComparar = siteComparar + 1
					var oddsIguais bool
					for i := siteComparar; i <= len(v.Sites)-1; i++ {
						oddsIguais = equal(x.Odds.H2H, v.Sites[i].Odds.H2H)
						if oddsIguais {
							break
						}

						go CalcularOddAnterior(oddAnterior)

						var proximoOdd float32
						if indexOdd == 0 {
							proximoOdd = v.Sites[i].Odds.H2H[1]
						} else {
							proximoOdd = v.Sites[i].Odds.H2H[0]
						}

						go CalcularProximaOdd(proximoOdd)
						s = s + fmt.Sprintf("%f :: ", proximoOdd)
					}

					if oddsIguais {
						break
					}
					siteComparar = siteindex
				}
				itemAnalisado.SitesIndex = append(itemAnalisado.SitesIndex, siteindex)
			}

			analisados = append(analisados, itemAnalisado)
		}

	}
	return sts
}

//CalcularOddAnterior Efetua o calculo de  viabilidade de  odd
func CalcularOddAnterior(odd float32) {

}

//CalcularProximaOdd Efetua o calculo de  viabilidade de  odd
func CalcularProximaOdd(odd float32) {

}

func equal(a []float32, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
