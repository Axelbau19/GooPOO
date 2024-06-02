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
	yearsExperience int
}
type HalfTimeEmployee struct {
	Person
	Employee
	internship bool
}

//Constructores

func NewFullTimeEmployee(name string, age int, id int, yearExperience int) *FullTimeEmployee {
	return &FullTimeEmployee{
		Person: Person{
			name: name,
			age:  age,
		},
		Employee: Employee{
			id: id,
		},
		yearsExperience: yearExperience,
	}
}

func (ftEmployee *FullTimeEmployee) GetMesagge() {
	fmt.Printf("FullTimeEmployee's Name :%s  with a age %d and him/her id is %d with %d years of experience\n", ftEmployee.name, ftEmployee.age, ftEmployee.id, ftEmployee.yearsExperience)
}

func NewHalfTimeEmployee(name string, age int, id int, internship bool) *HalfTimeEmployee {
	return &HalfTimeEmployee{
		Person: Person{
			name: name,
			age:  age,
		},
		Employee: Employee{
			id: id,
		},
		internship: internship,
	}
}

func (htEmployee *HalfTimeEmployee) GetMesagge() {
	fmt.Printf("HalfTimeEmployee's Name: %s, Age: %d, ID: %d, Internship: %t\n", htEmployee.name, htEmployee.age, htEmployee.id, htEmployee.internship)
}

type PrintInfo interface {
	GetMesagge()
}

func PrintData(p PrintInfo) {
	p.GetMesagge()
}

func main() {
	ft := NewFullTimeEmployee("Axel", 22, 1, 2)
	ht := NewHalfTimeEmployee("Alexis", 22, 2, true)
	PrintData(ft)
	PrintData(ht)
}
