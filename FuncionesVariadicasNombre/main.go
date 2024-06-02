package main

import "fmt"

//Funcion variadicas en donde puedo enviar numero de variables indeterminado
func suma(values ...int) int {
	//Creamos la variable
	total := 0
	//Ciclo donde vamos a obtener valores por que es un arreglo
	for _, num := range values {
		total += num
	}
	return total
}

//Creando una funcion donde reciba cadenas de texto indeterminado
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// para valores
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	//Go ya sabe que datos va regresar
	return
}
func main() {
	fmt.Println(suma(1))
	fmt.Println(suma(1, 2, 3, 4))
	fmt.Println(suma(1, 3, 4))
	printNames("Axel", "Alexis")
	fmt.Println(getValues(2))
}
