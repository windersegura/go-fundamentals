package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	separator := strings.Repeat("-", 20)
	fmt.Println(separator)

	for numeros := 0; numeros <= 3; numeros++ {
		fmt.Println(numeros)
	}

	fmt.Println(separator)

	for rango := range [3]int{} {
		fmt.Println(rango)
	}

	fmt.Println(separator)

	for {
		fmt.Println("Este bucle se ejecutará infinitamente.")
		break
	}

	for valor := range [6]int{} {
		if valor%2 == 0 {
			continue
		}
		fmt.Println(valor)
	}

}
