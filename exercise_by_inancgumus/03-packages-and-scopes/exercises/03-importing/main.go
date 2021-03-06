package main

// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

import (
	"fmt"
	a "fmt"
	b "fmt"
)

func main() {
	fmt.Println("hello")
	a.Println("hey")
	b.Println("hi")
}
