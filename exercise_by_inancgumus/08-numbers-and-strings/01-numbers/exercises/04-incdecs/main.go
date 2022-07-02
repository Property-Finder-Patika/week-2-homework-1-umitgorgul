package main

// ---------------------------------------------------------
// EXERCISE: Incdecs
//
//  1. Increase the `counter` 5 times
//  2. Decrease the `factor` 2 times
//  3. Print the product of counter and factor
//
// RESTRICTION
//  Use only the incdec statements
//
// EXPECTED OUTPUT
//  -75
// ---------------------------------------------------------

import "fmt"

func main() {
	counter, factor := 45, 0.5
	factor-- // - 0,5
	factor-- // - 1,5
	counter++ // 46
	counter++ // 47
	counter++ // 48
	counter++ // 49
	counter++ // 50
	fmt.Println(float64(counter) * factor) // -75
}
