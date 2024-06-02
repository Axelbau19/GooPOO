package main

import "fmt"

//Creacion de una clase
//Da valor por defecto
//Define propiedades
type Employee struct {
	id   int
	name string
}

//Crear funcion que posee la clase
func (e *Employee) setId(id int) {
	e.id = id
}

//Poner tu nombre variable y poner puntero en la clase
func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}
func (e *Employee) getName() string {
	return e.name
}

func main() {
	employee1 := Employee{}
	employee1.id = 0
	employee1.name = "Axel"
	fmt.Println("%v", employee1)
	employee1.setId(3)
	employee1.setName("Alexis")
	fmt.Println("%v", employee1)
	fmt.Println(employee1.getName())
	fmt.Println(employee1.getId())
}
