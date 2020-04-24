package main

import (
	"milhonarios/api"
	"milhonarios/models"
)

//FiltrarMaisDeUmSite filtra  mais de 1 site
func FiltrarMaisDeUmSite() []models.Odd {
	oddResponse1 := api.GetOddsFake("upcoming", "au")

	return oddResponse1.Filter(func(v models.Odd) bool {
		return v.SitesCount > 1
	})
}
