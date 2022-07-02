// An exercise for arithmetic operations
package main

import (
	"fmt"
)

func main() {

	num1 := 42
	num2 := 20

	// + add two variables
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)

	// - subtract two variables
	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, difference)

	// * multiply two variables
	multiply := num1 * num2
	fmt.Printf("%d * %d is %d\n", num1, num2, multiply)

	// / divide two integer variables
	division := num1 / num2
	fmt.Printf(" %d / %d = %d\n", num1, num2, division)

	num3 := 51.0
	num4 := 12.0

	// / divide two floating point variables
	result := num3 / num4
	fmt.Printf(" %g / %g = %g\n", num3, num4, result)

	// % modulo-divides two variables
	reminder := num1 % num2
	fmt.Println(reminder)

	num := 5

	num++
	fmt.Println(num) //6
	num--
	num--
	fmt.Println(num) // 4

	var result_1 int

	// assigns the value of num to result
	result_1 = num
	fmt.Println(result_1)

	// Compound Assignment Operators
	// += (addition assignment)
	// -= (subtraction assignment)
	// *= (multiplication assignment)
	// /= (division assignment)
	// %= (modulo assignment)
	number := 2
	number += 6

	// Relational Operator
	// a == b
	// a != b
	// a > b
	// a < b
	// a >= b
	// a <= b

	a := 5
	b := 6
	result_2 := a == b
	fmt.Println(result_2)

	// Logical Operator
	// exp1 && exp2 	*****	exp1 || exp2	***** 	!exp
}
