package main

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

import "fmt"

func main() {
	fmt.Printf("My name is %s and my lastname is %s.\n", "Umit", "Gorgul")

	// BONUS
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, "Umit", "Gorgul")
}
