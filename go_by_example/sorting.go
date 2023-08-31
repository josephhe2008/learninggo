package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("string:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	// isSorted return true false
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}