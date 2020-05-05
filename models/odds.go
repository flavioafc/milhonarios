package models

import (
	"math"
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
	Esporte     string
	Titulo      string
	Data        utils.Time
	Combinacoes []SitesCombinados
}

type SitesCombinados struct {
	Sites []ResultadoOdds
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

var percentualMaximo decimal.Decimal

//FilterSites filtra
func (vs *OddsResponse) FilterSites(aporteTotal float32, f func(criteria Odd) bool) []ResultadoFinal {
	var RetornoFinal []ResultadoFinal
	var resultadoFinal ResultadoFinal
	var sitesCombinados []SitesCombinados
	var titulo string

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

						var proximoOdd float32
						if indexOdd == 0 {
							proximoOdd = v.Sites[i].Odds.H2H[1]
						} else {
							proximoOdd = v.Sites[i].Odds.H2H[0]
						}

						var deOddMenor, ateOddMenor float32
						var deOddMaior, ateOddMaior float32

						var resultadoOddAnterior []ResultadoOdds
						var resultadoProximaOdd []ResultadoOdds

						if oddAnterior > proximoOdd {

							deOddMenor = (aporteTotal / proximoOdd) * 100 / aporteTotal
							deOddMaior = (aporteTotal / oddAnterior) * 100 / aporteTotal
							ateOddMenor = deOddMaior
							ateOddMaior = ateOddMenor + (ateOddMenor - deOddMenor) + 1

							if lucrativo(aporteTotal, int(deOddMaior), proximoOdd) {
								resultadoOddAnterior = CalcularOddAnterior(aporteTotal, oddAnterior, deOddMenor, ateOddMenor)
								resultadoProximaOdd = CalcularProximaOdd(aporteTotal, proximoOdd, deOddMaior, ateOddMaior)
							} else {
								continue
							}

						} else {

							deOddMenor = (aporteTotal / oddAnterior) * 100 / aporteTotal
							deOddMaior = (aporteTotal / proximoOdd) * 100 / aporteTotal
							ateOddMenor = deOddMenor + (deOddMenor - deOddMaior) + 1
							ateOddMaior = ateOddMenor

							if lucrativo(aporteTotal, int(deOddMaior), oddAnterior) {
								resultadoOddAnterior = CalcularOddAnterior(aporteTotal, oddAnterior, deOddMenor, ateOddMenor)
								resultadoProximaOdd = CalcularProximaOdd(aporteTotal, proximoOdd, deOddMaior, ateOddMaior)
							} else {
								continue
							}

						}
						var finalResultOdds []ResultadoOdds

						if resultadoOddAnterior != nil {
							if resultadoOddAnterior[0].Odd.GreaterThan(resultadoProximaOdd[0].Odd) {
								finalResultOdds = melhorLucratividadeEntreSites(aporteTotal, resultadoOddAnterior, resultadoProximaOdd, x.SiteNice, v.Sites[i].SiteNice)
							} else {
								finalResultOdds = melhorLucratividadeEntreSites(aporteTotal, resultadoProximaOdd, resultadoOddAnterior, v.Sites[i].SiteNice, x.SiteNice)
							}

							if titulo == "" || titulo == v.Teams[0]+" vs "+v.Teams[1] {
								// inserir nova combinação
								titulo = v.Teams[0] + " vs " + v.Teams[1]
								itensCombinados := finalResultOdds
								sitesCombinados = append(sitesCombinados, SitesCombinados{Sites: itensCombinados})
								resultadoFinal = ResultadoFinal{Esporte: v.SportNice, Titulo: titulo, Data: v.CommenceTime, Combinacoes: sitesCombinados}

							} else {
								RetornoFinal = append(RetornoFinal, resultadoFinal)

								titulo = v.Teams[0] + " vs " + v.Teams[1]
								itensCombinados := finalResultOdds
								sitesCombinados = append(sitesCombinados, SitesCombinados{Sites: itensCombinados})
								resultadoFinal = ResultadoFinal{Esporte: v.SportNice, Titulo: titulo, Data: v.CommenceTime, Combinacoes: sitesCombinados}
							}
						}
					}
					siteComparar = siteindex
				}
			}
		}
	}

	//se retornofinal tiver vazio e tiver algo em resultaFinal, então adiciona
	if len(RetornoFinal) == 0 && resultadoFinal.Titulo != "" {
		RetornoFinal = append(RetornoFinal, resultadoFinal)
	}
	return RetornoFinal
}

func lucrativo(aporteTotal float32, percentualdeMaiorOdd int, oddMenor float32) bool {
	valor1 := aporteTotal * float32(percentualdeMaiorOdd) / 100
	valor2 := aporteTotal - valor1
	resultado := valor2 * oddMenor

	if resultado <= aporteTotal {
		return false
	}

	return true
}

func melhorLucratividadeEntreSites(aporteTotal float32, maiorResultadoOdd []ResultadoOdds, menorResultadoOdd []ResultadoOdds, siteAnterior string, proximoSite string) []ResultadoOdds {
	var resultado []ResultadoOdds

	j := len(menorResultadoOdd) - 1
	for i := 0; i <= len(maiorResultadoOdd); i++ {
		subtracao := menorResultadoOdd[j].Lucro.Sub(maiorResultadoOdd[i].Lucro)
		if subtracao.LessThan(decimal.Zero) {
			maiorResultadoOdd[i-1].Site = siteAnterior
			menorResultadoOdd[j+1].Site = proximoSite
			resultado = append(resultado, maiorResultadoOdd[i-1], menorResultadoOdd[j+1])
			break
		}
		j--
	}

	return resultado
}

//CalcularOddAnterior Efetua o calculo de  viabilidade de  odd
func CalcularOddAnterior(aporteTotal float32, odd float32, de float32, ate float32) []ResultadoOdds {
	return calculo(aporteTotal, odd, de, ate)
}

//CalcularProximaOdd Efetua o calculo de  viabilidade de  odd
func CalcularProximaOdd(aporteTotal float32, odd float32, de float32, ate float32) []ResultadoOdds {
	return calculo(aporteTotal, odd, de, ate)
}

func calculo(aporteTotal float32, odd float32, de float32, ate float32) []ResultadoOdds {
	var responseResultado []ResultadoOdds

	//converte para 2 casas decimais
	de = float32(math.Round(float64(de)))

	var i float32
	for i = de; i <= ate; i++ {
		percentual := i / 100

		aporteInvestido := aporteTotal * percentual
		lucro := odd * aporteInvestido

		itemResultado := ResultadoOdds{Odd: decimal.NewFromFloat32(odd), Percentual: float32(i), AporteInvestido: decimal.NewFromFloat32(aporteInvestido), Lucro: decimal.NewFromFloat32(lucro)}
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
