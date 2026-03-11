package main

import "fmt"

func modificarArreglo(arg *[6]int) {
	(*arg)[0] = 100
}

func main() {
	arreglo := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Arreglo antes de modificar: ", arreglo)
	modificarArreglo(&arreglo)
	fmt.Println("Arreglo despues de modificar: ", arreglo)

}
