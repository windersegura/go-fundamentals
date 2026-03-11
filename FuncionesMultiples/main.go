package main

import "fmt"

func nombresMultiples(nombre1 string, nombre2 string, apellido string) (map[string]string, map[string]string) {
	mapa1 := make(map[string]string)
	mapa2 := make(map[string]string)

	mapa1[nombre1] = apellido
	mapa2[nombre2] = apellido

	return mapa1, mapa2
}

func main() {
	nombre1 := "Winder"
	nombre2 := "Sleam"
	apellido := "Segura"

	mapa1, mapa2 := nombresMultiples(nombre1, nombre2, apellido)

	fmt.Println("mapa1: ", mapa1)
	fmt.Println("mapa2: ", mapa2)

	_, dato := nombresMultiples(nombre1, nombre2, apellido)
	fmt.Println("dato: ", dato)

}
