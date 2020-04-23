package main

import (
	"fmt"
	"milhonarios/api"
)

func main() {

	oddResponse := api.GetOddsFake("upcoming", "eu")
	fmt.Println(len(oddResponse.Data))
	fmt.Println(oddResponse.Data)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(sportResponse.Data[i].Key + " * " + sportResponse.Data[i].Title)
	// }

}
