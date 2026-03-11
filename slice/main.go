package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string

	s = make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("The slice is: ", s)

	s = append(s, "d", "e")
	fmt.Println("The slice is now: ", s)

	s2 := []string{"f", "g", "h"}
	fmt.Println("The second slice is: ", s2)

	s3 := []string{"f", "g", "h"}
	fmt.Println(slices.Equal(s2, s3))

	for i := 0; i < len(s2); i++ {
		s = append(s, s2[i])
	}

	fmt.Println("The slice is now: ", s)

}
