package main

import "fmt"

/// chan<- sirve para decir que es de escritura
// <- chan sirve solamente para lectura
func Generator(chanel chan<- int) {
	for i := 0; i <= 10; i++ {
		chanel <- i
	}
	close(chanel)
}

func Double(input <-chan int, output chan<- int) {
	for value := range input {
		output <- 2 * value
	}
	close(output)
}

func Print(chanel <-chan int) {
	for value := range chanel {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)
	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
