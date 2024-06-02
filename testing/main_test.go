package main

import "testing"

func TestSuma(t *testing.T) {
	/*total := Suma(10, 5)
	if total != 7 {
		t.Errorf("La suma es incorrecta , tiene %d y tenia como expectativa %d", total, 7)
	}*/
	//Test de tablas
	tablesCases := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{25, 26, 51},
	}
	for _, item := range tablesCases {
		total := Suma(item.a, item.b)
		if total != item.n {
			t.Errorf("La suma fue incorrecta, el resultado es %d y se esperaba %d ", total, item.n)
		}
	}
}

//Go tiene su propia libreria para test

//Creamos un archivo con el comando go test -coverprofile=nombreArchivo.out

func TestMax(t *testing.T) {
	tablesCases := []struct {
		a   int
		b   int
		max int
	}{
		{4, 2, 4},
		{2, 1, 2},
		{1, 2, 2},
	}
	for _, item := range tablesCases {
		maxNumber := GetMax(item.a, item.b)
		if maxNumber != item.max {
			t.Errorf("El numero %d esta mal, el mayor deber ser %d", maxNumber, item.max)
		}
	}
}

//Fibonacci
func TestFibonacci(t *testing.T) {
	tableCase := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tableCase {
		fib := Fibonnaci(item.a)
		if fib != item.n {
			t.Errorf("Error en metodo fibonnaci")
		}
	}
}
