package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	fmt.Println("Hello, World!")
	time.Sleep(2 * time.Second)
	fmt.Println("Goodbye, World!")
	c <- true // Envía un valor al canal para indicar que la goroutine ha terminado
}

func main() {
	c := make(chan bool) // Crea un canal para sincronización
	go hello(c)          // Ejecuta la función hello() en una goroutine

	fmt.Println("This is the main function.")

	<-c // Espera a recibir un valor del canal, lo que indica que la goroutine ha terminado
	fmt.Println("The goroutine has finished.")
}
