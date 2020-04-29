package main

import (
	"milhonarios/models"
)

//FiltrarMaisDeUmSite filtra  mais de 1 site
func FiltrarMaisDeUmSite(aporteTotal float32, odds models.OddsResponse) []models.ResultadoFinal {
	return odds.FilterSites(aporteTotal, func(v models.Odd) bool {
		return v.SitesCount > 1
	})
}

func equal(a, b []float32) bool {
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
