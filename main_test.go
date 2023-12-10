package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 2)

	if result != 4 {
		t.Error("Erro ao somar 2 + 2, o reultado deveria ser 4")
	}
}

func TestSub(t *testing.T) {
	result := sub(3, 2)

	if result != 1 {
		t.Error("Erro ao subtrair 3 - 2, o reultado deveria ser 1")
	}
}

func TestTimes(t *testing.T) {
	result := time(2, 2)

	if result != 4 {
		t.Error("Erro ao multiplicar 2 * 2, o reultado deveria ser 4")
	}
}

func TestMasterSum(t *testing.T) {
	result := masterSum(2, 2)

	if result != 6 {
		t.Error("Erro ao somar 2 + 2 + 2, o reultado deveria ser 6")
	}
}

func TestMain(t *testing.T) {
	main()
}
