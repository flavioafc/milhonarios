package main

import (
	"fmt"
	"milhonarios/api"
	"strconv"
)

func main() {

	sportResponse := api.GetSports()

	for i := 0; i < len(sportResponse.Data); i++ {
		fmt.Println(sportResponse.Data[i].Key)
	}

	fmt.Printf("* Serão " + strconv.Itoa(len(sportResponse.Data)) + " chamadas!")
}
