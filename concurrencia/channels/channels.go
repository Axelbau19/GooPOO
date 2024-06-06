package main

import "fmt"

func main() {
	///////////////////////Aqui se ponea la capacidad del canal
	chanel := make(chan int, 3)

	chanel <- 1
	chanel <- 2

	fmt.Println(<-chanel)
	fmt.Println(<-chanel)
}
