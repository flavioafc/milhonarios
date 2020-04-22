package main

import (
	"fmt"
	"milhonarios/api"
)

func main() {

	oddResponse := api.GetOddsFake("uk")
	fmt.Println(oddResponse)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(sportResponse.Data[i].Key + " * " + sportResponse.Data[i].Title)
	// }

}
