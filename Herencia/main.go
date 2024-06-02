package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Composicion de herencia, se agrega el tipo de objeto
type FullTimeEmployee struct {
	Person
	Employee
}

//Metodo

func NewFullTimeEmployee(name string, age int, id int) *FullTimeEmployee {
	return &FullTimeEmployee{
		Person: Person{
			name: name,
			age:  age,
		},
		Employee: Employee{
			id: id,
		},
	}
}

func (ftEmployee *FullTimeEmployee) GetMesagge() {
	fmt.Printf("FullTimeEmployee's Name :%s  with a age %d and him/her id is %d\n", ftEmployee.name, ftEmployee.age, ftEmployee.id)
}
func main() {
	ft := NewFullTimeEmployee("Axel", 22, 1)
	ft.GetMesagge()
}
