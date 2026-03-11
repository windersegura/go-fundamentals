package main

import "fmt"

func main() {
	edad := 18

	if edad >= 18 {
		fmt.Println("Eres mayor de edad.")
	} else {
		fmt.Println("Eeres menor de edad.")
	}

	if 8%2 == 0 {
		fmt.Println("El numero es par.")
	} else {
		fmt.Println("El numero es impar.")
	}

	if numero := 9; numero < 0 {
		fmt.Println("El numero es negativo")
	} else if numero < 10 {
		fmt.Println("El numero es de un solo digito")
	} else {
		fmt.Println("El numero es mayor a nueve")
	}
}
