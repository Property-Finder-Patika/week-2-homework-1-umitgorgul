package main

// ---------------------------------------------------------
// EXERCISE: Area
//
//  Fix the following program.
//
// RESTRICTION
//  You should not use any variables.
//
// EXPECTED OUTPUT
//  area = 1250
// ---------------------------------------------------------

import "fmt"

func main() {
	// const (
	// 	width  int16 = 25
	// 	height int32 = width * 2
	// )

	// fmt.Printf("area = %d\n", width*height)

	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}
