package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		channel1 <- "Finaliza primera Routine"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		channel2 <- "Finaliza segunda Routine"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		}
	}
}
