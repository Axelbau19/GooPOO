package main

import "testing"

// Creamos los mocks
func TestGetFulltimeEmployeeById(t *testing.T) {

	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1, // Usar `Id` en lugar de `id`
						Position: "CEO",
					}, nil
				}
				GetPersonByDni = func(dni string) (Person, error) {
					return Person{
						Name: "Axel",
						Age:  30,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  22,
					DNI:  "1",
					Name: "Axel",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDni

	for _, test := range table {
		test.mockFunc()
		ftEmployee, err := GetFullTimeEmployee(test.id, test.dni)
		if err != nil {
			t.Errorf("Error inesperado: %v", err)
		}
		if ftEmployee.Age != test.expectedEmployee.Age {
			t.Errorf("Error, se tiene %d y era %d", ftEmployee.Age, test.expectedEmployee.Age)
		}
	}
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDni = originalGetPersonByDni
}
