package main

import (
	"fmt"
	"milhonarios/api"
)

func main() {

	oddResponse1 := api.GetOddsFake("upcoming", "us_domingo")
	filtrado1 := FiltrarMaisDeUmSite(500, oddResponse1)

	oddResponse2 := api.GetOddsFake("upcoming", "us")
	filtrado2 := FiltrarMaisDeUmSite(500, oddResponse2)

	oddResponse3 := api.GetOddsFake("upcoming", "us_2")
	filtrado3 := FiltrarMaisDeUmSite(500, oddResponse3)

	oddResponseEuropa := api.GetOddsFake("upcoming", "eu")
	filtradoEuropa := FiltrarMaisDeUmSite(500, oddResponseEuropa)

	fmt.Println("---------------------------")
	fmt.Println(filtrado1)

	fmt.Println("---------------------------")
	fmt.Println(filtrado2)

	fmt.Println("---------------------------")
	fmt.Println(filtrado3)

	fmt.Println("---------------------------")
	fmt.Println(filtradoEuropa)
}
