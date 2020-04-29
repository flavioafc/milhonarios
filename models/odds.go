package models

import (
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

//ResultadoFinal exibe os resultados pro front
type ResultadoFinal struct {
	Esporte string
	Titulo  string
	Data    utils.Time
	Sites   []ResultadoOdds
}

//ResultadoOdds exibe o resultado final
type ResultadoOdds struct {
	Site            string
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
func (vs *OddsResponse) FilterSites(aporteTotal float32, f func(criteria Odd) bool) []ResultadoFinal {
	var RetornoFinal []ResultadoFinal

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
							continue
						}

						resultadoOddAnterior := CalcularOddAnterior(aporteTotal, oddAnterior)

						var proximoOdd float32
						if indexOdd == 0 {
							proximoOdd = v.Sites[i].Odds.H2H[1]
						} else {
							proximoOdd = v.Sites[i].Odds.H2H[0]
						}

						resultadoProximaOdd := CalcularProximaOdd(aporteTotal, proximoOdd)

						var finalResultOdds []ResultadoOdds
						if resultadoOddAnterior[0].Odd.GreaterThan(resultadoProximaOdd[0].Odd) {
							finalResultOdds = comparar(aporteTotal, resultadoOddAnterior, resultadoProximaOdd, x.SiteNice, v.Sites[i].SiteNice)
						} else {
							finalResultOdds = comparar(aporteTotal, resultadoProximaOdd, resultadoOddAnterior, v.Sites[i].SiteNice, x.SiteNice)
						}

						resultadoFinal := ResultadoFinal{Esporte: v.SportNice, Titulo: v.Teams[0] + " vs " + v.Teams[1], Data: v.CommenceTime, Sites: finalResultOdds}
						RetornoFinal = append(RetornoFinal, resultadoFinal)
					}
					siteComparar = siteindex
				}
			}
		}
	}
	return RetornoFinal
}

func comparar(aporteTotal float32, maiorResultadoOdd []ResultadoOdds, menorResultadoOdd []ResultadoOdds, siteAnterior string, proximoSite string) []ResultadoOdds {

	maiorResultOdd := maiorResultadoOdd[len(maiorResultadoOdd)-1]
	indiceMaiorLucroMenorOdd := len(menorResultadoOdd) - 1
	var resultado []ResultadoOdds

	for i := indiceMaiorLucroMenorOdd; i >= 0; i-- {
		somatorio := maiorResultOdd.AporteInvestido.Add(menorResultadoOdd[i].AporteInvestido)
		aporteTotal := decimal.NewFromFloat32(aporteTotal)

		if somatorio.Equal(aporteTotal) {
			maiorResultOdd.Site = siteAnterior
			menorResultadoOdd[i].Site = proximoSite

			resultado = append(resultado, maiorResultOdd, menorResultadoOdd[i])
		}
	}

	return resultado
}

//CalcularOddAnterior Efetua o calculo de  viabilidade de  odd
func CalcularOddAnterior(aporteTotal float32, odd float32) []ResultadoOdds {
	return calculo(aporteTotal, odd)
}

//CalcularProximaOdd Efetua o calculo de  viabilidade de  odd
func CalcularProximaOdd(aporteTotal float32, odd float32) []ResultadoOdds {
	return calculo(aporteTotal, odd)
}

func calculo(aporteTotal float32, odd float32) []ResultadoOdds {
	var responseResultado []ResultadoOdds
	valorAporteTotal := decimal.NewFromFloat32(aporteTotal)
	valorOdd := decimal.NewFromFloat32(odd)

	var i float32
	for i = 1; i <= 100; i++ {
		percentual := decimal.NewFromFloat32(i / 100)

		aporteInvestido := valorAporteTotal.Mul(percentual)
		lucro := valorOdd.Mul(aporteInvestido)

		itemResultado := ResultadoOdds{Odd: valorOdd, Percentual: i, AporteInvestido: aporteInvestido, Lucro: lucro}
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
