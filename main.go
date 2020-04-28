package main

import (
	"fmt"
	"milhonarios/api"
)

func main() {

	oddResponse1 := api.GetOddsFake("upcoming", "us_2")
	filtrado := FiltrarMaisDeUmSite(oddResponse1)

	fmt.Println(filtrado)
	//odds := FiltrarMaisDeUmSite(oddResponse1)

	//OddsDiferentes(odds)

}
