package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	DOB       string
}

func main() {

	joseph := Person{"Joe", "H", 42, "1981-01-01"}
	fmt.Println(joseph)

}