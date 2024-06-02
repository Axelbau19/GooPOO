package main

import "fmt"

//Se puede crear diferentes formas de constructor pero me gusto mas este

//Definir la clase
type Car struct {
	marca string
	anio  int
}

//Constructor por metodo
func NewCar(marca string, anio int) *Car {
	return &Car{
		marca: marca,
		anio:  anio,
	}
}

func main() {
	e := NewCar("BWM", 2021)
	fmt.Println(*e)
}
