package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Decalre and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Slice of strings
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	// Print length of slice
	fmt.Println(len(fruitSlice))

	// Print a slice of the slice
	fmt.Println(fruitSlice[1:3])

	// Create a slice with make
	fruitSlice2 := make([]string, 4)
	fruitSlice2[0] = "Apple"
	fruitSlice2[1] = "Orange"
	fruitSlice2[2] = "Grape"
	fruitSlice2[3] = "Cherry"

	// Print the slice
	fmt.Println(fruitSlice2)

	// Print a slice of the slice
	fmt.Println(fruitSlice2[1:3])
}
