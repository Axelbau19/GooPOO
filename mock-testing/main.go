package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDni = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	//SELECT * FROM PERSONA WHERE...
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil

}

var GetFullTimeEmployee = func(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e
	p, err := GetPersonByDni(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p
	return ftEmployee, nil
}

///Mockear servicio
