package main

import (
    "fmt"
    "os"
)

func main() {

	// defer won't run
    defer fmt.Println("!")

	// os.Exit exits immediately
    os.Exit(3)
}