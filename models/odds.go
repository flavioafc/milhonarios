package models

import (
	"fmt"
	"milhonarios/utils"

	"github.com/shopspring/decimal"
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

//Resultado exibe o resultado final
type Resultado struct {
	Odd             decimal.Decimal
	Percentual      float32
	AporteInvestido decimal.Decimal
	Lucro           decimal.Decimal
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

//FilterSites filtra
func (vs *OddsResponse) FilterSites(f func(criteria Odd) bool) []Sites {
	sts := make([]Sites, 0)
	for _, v := range vs.Data {
		if f(v) {
			for siteindex, x := range v.Sites {
				sts = append(sts, x)
				siteComparar := siteindex

				for indexOdd, oddAnterior := range x.Odds.H2H {
					//comparar com  o proximo Odd do proximo site
					siteComparar = siteComparar + 1
					var oddsIguais bool
					for i := siteComparar; i <= len(v.Sites)-1; i++ {
						oddsIguais = equal(x.Odds.H2H, v.Sites[i].Odds.H2H)
						if oddsIguais {
							break
						}
						resultadoOddAnterior := CalcularOddAnterior(500, oddAnterior)

						var proximoOdd float32
						if indexOdd == 0 {
							proximoOdd = v.Sites[i].Odds.H2H[1]
						} else {
							proximoOdd = v.Sites[i].Odds.H2H[0]
						}

						resultadoProximaOdd := CalcularProximaOdd(500, proximoOdd)

						fmt.Println(resultadoOddAnterior)
						fmt.Println("----------------------------------------------------------------")
						fmt.Println(resultadoProximaOdd)
						fmt.Println("============================= FIM  FIM  FIM  FIM ====================================")
					}

					if oddsIguais {
						break
					}
					siteComparar = siteindex
				}
			}
		}

	}
	return sts
}

//CalcularOddAnterior Efetua o calculo de  viabilidade de  odd
func CalcularOddAnterior(aporteTotal float32, odd float32) []Resultado {
	return calculo(aporteTotal, odd)
}

//CalcularProximaOdd Efetua o calculo de  viabilidade de  odd
func CalcularProximaOdd(aporteTotal float32, odd float32) []Resultado {
	return calculo(aporteTotal, odd)
}

func calculo(aporteTotalSite float32, odd float32) []Resultado {
	var responseResultado []Resultado
	valorAporteTotal := decimal.NewFromFloat32(aporteTotalSite)
	valorOdd := decimal.NewFromFloat32(odd)

	var i float32
	for i = 1; i <= 100; i++ {
		percentual := decimal.NewFromFloat32(i / 100)

		aporteInvestido := valorAporteTotal.Mul(percentual)
		lucro := valorOdd.Mul(aporteInvestido)

		itemResultado := Resultado{Odd: valorOdd, Percentual: i, AporteInvestido: aporteInvestido, Lucro: lucro}
		if lucro.GreaterThan(valorAporteTotal) {
			responseResultado = append(responseResultado, itemResultado)
			break
		}
		responseResultado = append(responseResultado, itemResultado)
	}

	return responseResultado
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
