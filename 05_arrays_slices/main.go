package main

import (
	"fmt"
)

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

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

	fruitSlice2 := make([]string, 4)
	fruitSlice2[0] = "Apple"
	fruitSlice2[1] = "Orange"
	fruitSlice2[2] = "Grape"
	fruitSlice2[3] = "Cherry"

	fmt.Println(fruitSlice2)
	fmt.Println(fruitSlice2[1:3])

}
