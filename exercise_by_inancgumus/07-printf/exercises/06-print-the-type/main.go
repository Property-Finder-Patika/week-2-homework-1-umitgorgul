package main

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

import "fmt"

func main() {
	// %T reads arg #2, but call has 1 arg so we use %[1]T
	fmt.Printf("Type of %d is %[1]T\n", 3)
}
