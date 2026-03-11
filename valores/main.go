package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero int = 10
	numero2 := 20.5
	fmt.Println(numero, numero2)

	var numeroentero = 10
	numerodoble := 20.5
	resultado := numeroentero + int(numerodoble)
	fmt.Println(resultado)

	var nombre string = "Winder"
	apellido := "Segura"

	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	var numeroString string = strconv.Itoa(numero)
	fmt.Println(numeroString)

	numeroConvertido, err := strconv.Atoi(numeroString)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println(numeroConvertido)
	}
}
