package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Refactor Feet to Meter
//
//  Remember the feet to meters program?
//  Now, it's time to refactor it.
//  Define your own Feet and Meters types.
//
//  Follow the steps inside the code.
// ---------------------------------------------------------

func main() {
	// ----------------------------
	// 1. Define Feet and Meters types below
	//    Their underlying type can be float64
	// ...

	// ----------------------------
	// 2. UNCOMMENT THE CODE BELOW THEN DON'T TOUCH IT
	// var (
	// 	feet   Feet
	// 	meters Meters
	// )

	// ----------------------------
	// 3. Get feet value from the command-line
	// 4. Convert it to an float64 first using ParseFloat
	// 5. Then, convert it into a Feet type
	// ... TYPE YOUR CODE HERE

	// ----------------------------
	// 6. Uncomment the code below
	// 7. And, convert the expression to Meters type

	// meters = feet * 0.3048

	// ----------------------------
	// 8. UNCOMMENT THE CODE BELOW
	// fmt.Printf("%g feet is %g meters.\n", feet, meters)

	type (
		Feet   float64
		Meters float64
	)

	var (
		feet   Feet
		meters Meters
	)

	arg := os.Args[1]

	// feet is a Feet value now
	val, _ := strconv.ParseFloat(arg, 64)
	feet = Feet(val)

	// meters is a Meters value now
	meters = Meters(feet * 0.3048)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
