package main

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

import (
	"fmt"
	"os"
)

func main() {
	var (
		l = len(os.Args) - 1
		a = os.Args[1]
		b = os.Args[2]
		c = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", a, "!")
	fmt.Println("Hello great", b, "!")
	fmt.Println("Hello great", c, "!")
	fmt.Println("Nice to meet you all.")
}
