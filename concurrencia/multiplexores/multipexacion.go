package main

import (
	"fmt"
	"time"
)

// /Conexion entre canales
func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	duracion := 4 * time.Second
	duracion2 := 2 * time.Second
	go DoSomething(duracion, channel1, 1)
	go DoSomething(duracion2, channel2, 2)
	for i := 0; i < 2; i++ {
		//Switch
		select {
		case channelMsg1 := <-channel1:
			fmt.Println(channelMsg1)
		case channel2Msg := <-channel2:
			fmt.Println(channel2Msg)
		}
	}
}

func DoSomething(i time.Duration, channel chan<- int, param int) {
	time.Sleep(i)
	channel <- param
}
