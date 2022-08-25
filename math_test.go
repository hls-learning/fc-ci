package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resutado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtrai(t *testing.T) {
	total := Subtrai(15, 15)

	if total != 0 {
		t.Errorf("Resutado da soma é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}

func TestMultiplica(t *testing.T) {
	total := Multiplica(15, 15)

	if total != 225 {
		t.Errorf("Resutado da soma é inválido: Resultado %d. Esperado: %d", total, 225)
	}
}

func TestDivide(t *testing.T) {
	total := Divide(15, 15)

	if total != 1 {
		t.Errorf("Resutado da soma é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}
