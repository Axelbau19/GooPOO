package main

import (
	"fmt"
	"time"
)

func main() {
	/*	x := 5
		//creamos una funcion anonima
		y := func() int {
			return x * 10
		}()
		fmt.Println(y)
		z := 10
		w := func() int {
			return z * 2
		}()
		fmt.Print(w)*/
	c := make(chan int)
	go func() {
		fmt.Println("Inicia la funcion")
		time.Sleep(5 * time.Second)
		fmt.Println("Finalizo la go routine")
		c <- 1
	}()
	<-c

}
