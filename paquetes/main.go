package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("MY_ENV_VAR")

	if envVar == "" {
		fmt.Println("The environment variable MY_ENV_VAR is not set.")
	} else {
		fmt.Println("The value of MY_ENV_VAR is:", envVar)
	}

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()
	fmt.Println("File example.txt created successfully.")
}
