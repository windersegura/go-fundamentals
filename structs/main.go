package main

import "fmt"

type Persona struct {
	Nombre    string
	Edad      int
	Profesion string
}

func nuevaPersona(nombre string, edad int, profesion string) *Persona {
	nuevaPersona := Persona{
		Nombre:    nombre,
		Edad:      edad,
		Profesion: profesion,
	}
	return &nuevaPersona

}

func main() {
	p := nuevaPersona("Winder", 27, "Ingeniero")
	fmt.Println(*p)

	personita := Persona{"Winder", 28, "Devops"}
	fmt.Println(personita.Nombre)

	personidta2 := &personita
	fmt.Println(personidta2.Profesion)

}
