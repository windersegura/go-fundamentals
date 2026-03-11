package main

import "fmt"

func main() {
	mapa := make(map[string]int)

	mapa["Winder"] = 4
	mapa["Segura"] = 8
	fmt.Println("mapa: ", mapa)

	version1 := mapa["Winder"]
	fmt.Println("version1: ", version1)

	version2 := mapa["Segura"]
	fmt.Println("version2: ", version2)

	fmt.Println("longitud del mapa: ", len(mapa))

	_, dato := mapa["Winder"]
	fmt.Println("dato: ", dato)

	delete(mapa, "Winder")
	fmt.Println("mapa después de borrar Winder: ", mapa)
}
