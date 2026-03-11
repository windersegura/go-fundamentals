package main

import (
	"fmt"
	"time"
)

func main() {
	dias := [...]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}
	fmt.Println(dias)
	fmt.Println("Tamaño de arreglo:", len(dias))

	t := time.Now().Weekday()

	switch t {
	case time.Monday:
		fmt.Println(dias[0])
	case time.Tuesday:
		fmt.Println(dias[1])
	case time.Wednesday:
		fmt.Println(dias[2])
	case time.Thursday:
		fmt.Println(dias[3])
	case time.Friday:
		fmt.Println(dias[4])
	case time.Saturday:
		fmt.Println(dias[5])
	case time.Sunday:
		fmt.Println(dias[6])
	}
}
