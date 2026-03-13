package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Resultado goroutine 1"
	}()

	select {
	case result := <-c1:
		fmt.Println("Resultado: ", result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Resultado goroutine 2"
	}()

	select {
	case result := <-c2:
		fmt.Println("Resultado: ", result)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout")
	}

}
