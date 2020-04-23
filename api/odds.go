package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"milhonarios/models"
	"net/http"
	"os"
)

//GetSports chama api para consumir dados de todos os sports
func GetSports() models.SportsResponse {
	var sportsResponse models.SportsResponse
	response, err := http.Get("http://api.the-odds-api.com/v3/sports/?apiKey=f2be2d2d006a74f6dccb2faa7aff2a97&all=1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &sportsResponse)

	return sportsResponse
}

//GetOdds obtem os odds da api
func GetOdds(sportkey string, region string) models.OddsResponse {
	var oddsResponse models.OddsResponse
	url := fmt.Sprintf("https://api.the-odds-api.com/v3/odds/?sport=%s&region=%s&mkt=h2h&apiKey=f2be2d2d006a74f6dccb2faa7aff2a97", sportkey, region)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &oddsResponse)

	return oddsResponse
}

//GetOddsFake é um  fake para testar a estrutura
func GetOddsFake(sportkey string, region string) models.OddsResponse {
	var oddsResponse models.OddsResponse
	// read file
	data, err := ioutil.ReadFile("./mocks/odds_" + sportkey + "_" + region + ".json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &oddsResponse)
	if err != nil {
		fmt.Println("error:", err)
	}

	return oddsResponse
}
