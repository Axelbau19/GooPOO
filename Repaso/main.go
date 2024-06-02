package main

import (
	"fmt"
	"strconv"
	"time"
)

// Creacion de una goroutine
func goRoutine(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Go routine")
	c <- 1
}

func main() {
	//Variable explicita
	var x int
	x = 2
	//Variable implicita
	y := 3
	fmt.Println(x, y)
	//go build main para nivel a produccion
	//En go no existe el try catch
	//se tiene que decir de forma explicita
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	//Declarar una excepcion
	if err != nil {
		fmt.Println("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
	//Maps
	m := make(map[string]int)
	m["key"] = 10

	fmt.Println(m["key"])
	//Slice
	s := []int{1, 2, 3}
	for index, valu := range s {
		fmt.Println(index, valu)
	}
	s = append(s, 12)
	for index, valu := range s {
		fmt.Println(index, valu)
	}
	//Creamos el canal (se necesita para un go routine un canal)
	//c := make(chan int)

	//go goRoutine(c)
	//<-c
	//Apuntador
	g := 25
	h := &g
	fmt.Println(*h)

}
