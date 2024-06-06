package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup, chanel chan int) {
	defer wg.Done()
	fmt.Printf("Id %d inicio\n", i)
	time.Sleep(5 * time.Second)
	fmt.Printf("Id %d Finalizo\n", i)
	<-chanel
}

func main() {
	channel := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		channel <- 1
		wg.Add(1)
		go doSomething(i, &wg, channel)
	}
	wg.Wait()
}

//Espacios y procesos que van utilizar
