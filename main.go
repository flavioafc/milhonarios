package main

import (
	"fmt"
	"milhonarios/api"
)

func main() {

	oddResponse1 := api.GetOddsFake("upcoming", "us_domingo")
	filtrado := FiltrarMaisDeUmSite(500, oddResponse1)

	fmt.Println(filtrado)
}
