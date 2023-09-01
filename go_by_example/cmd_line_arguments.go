package main

import (
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]
	arg0 := os.Args[0]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
	// arg0 is the program
	fmt.Println(arg0)
}