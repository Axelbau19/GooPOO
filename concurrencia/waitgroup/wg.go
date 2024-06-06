package main

import (
	"fmt"
	"sync"
	"time"
)

func prueba(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Inicio %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizo")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go prueba(i, &wg)
	}
	wg.Wait()
}
