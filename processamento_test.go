package main

import (
	"testing"
)

func Test_Deve_Retornar_Jogos_Com_Mais_de_Um_site(t *testing.T) {
	x := FiltrarMaisDeUmSite()

	result := len(x)
	if result == 1 {
		t.Errorf("Deve retornar mais de 1. Resultado = %d", result)
	} else {
		t.Logf("Passou, trouxe %d resultados", result)
	}
}
