package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {

	person := Person{
		Name:  "Winder",
		Age:   27,
		Email: "segurawinder@gmail.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error al codificar a JSON: ", err)
	}

	fmt.Println("El objeto en JSON es: ", string(jsonData))

	/// Convertir de Json a struct

	jsonString := `{"name":"Winder","age":27,"email":"segurawinder@gmail.com"}`
	var person2 Person
	err = json.Unmarshal([]byte(jsonString), &person2)
	if err != nil {
		fmt.Println("Error al decodificar JSON: ", err)
	}
	fmt.Println("El objeto decodificado es: ", person2)

}
