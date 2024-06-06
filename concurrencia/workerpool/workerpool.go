package main

import "fmt"

func Worker(id int, jobs <-chan int, resultado chan<- int) {
	for job := range jobs {
		fmt.Printf("Trabajo con id  %d inicio con %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Trabajo con id %d, trabajo %d y fib %d\n", id, job, fib)
		resultado <- fib
	}
}

func Fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}

//Se divide para calcular el fibonaci
func main() {
	task := []int{2, 3, 4, 5, 7, 10, 40}
	nWork := 3
	channelJobs := make(chan int, len(task))
	resultadoChannel := make(chan int, len(task))
	for i := 0; i < nWork; i++ {
		go Worker(i, channelJobs, resultadoChannel)
	}
	for _, value := range task {
		channelJobs <- value
	}
	close(channelJobs)

	for j := 0; j < len(task); j++ {
		<-resultadoChannel
	}

}
