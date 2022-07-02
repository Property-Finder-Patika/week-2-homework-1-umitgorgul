package main

// ---------------------------------------------------------
// EXERCISE: Redeclare
//
// 	1. Short declare two int variables: age and yourAge
//     (use multiple short declaration syntax)
//
//  2. Short declare a new float variable: ratio
//     And, change the 'age' variable to 42
//
//     (! You should use redeclaration)
//
//  4. Print all the variables
//
// EXPECTED OUTPUT
//  42, 20, 3.14
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	age, yourAge := 1, 20
	age, ratio := 42, 3.14

	fmt.Println(age, yourAge, ratio)
}
