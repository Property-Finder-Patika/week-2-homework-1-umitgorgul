package main

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)
}
