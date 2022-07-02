// An exercise for conversions among basic types
package main

import (
	"fmt"
)

func main() {
	var a bool = false      // Boolean
	var b int = 8           // Integer
	var c float32 = 9.99    // Floating point number
	var d string = "Selam!" // String

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	// uint keywords, can only store non-negative values
	var x uint = 500
	var y uint = 4500

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}
