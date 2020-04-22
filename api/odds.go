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

//Sports chama api para consumir dados de todos os sports
func Sports() models.Response {
	var responseSport models.Response
	response, err := http.Get("http://api.the-odds-api.com/v3/sports/?apiKey=f2be2d2d006a74f6dccb2faa7aff2a97&all=1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseSport)

	return responseSport
}
