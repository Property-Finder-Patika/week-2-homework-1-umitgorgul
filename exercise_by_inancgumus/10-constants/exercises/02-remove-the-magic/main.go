package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Remove the Magic
//
//  Get rid of the magic numbers in the following code.
//
// RESTRICTIONS
//  1. You should declare 3 constants named:
//       hoursInDay, daysInWeek, and hoursPerWeek
//
//  2. And, hoursPerWeek constant should be initialized
//     using hoursInDay and daysInWeek constants.
//
// EXPECTED OUTPUT
//  There are 840 hours in 5 weeks.
// ---------------------------------------------------------
func main() {
	const (
		totalHours       = 24
		totalDays        = 7
		totalHoursInWeek = totalHours * totalDays
	)

	fmt.Printf("There are %d hours in %d weeks.\n",
		totalHoursInWeek*5, 5)
}
