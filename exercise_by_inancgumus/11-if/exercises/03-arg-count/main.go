package main

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------
import (
	"fmt"
	"os"
)

func main() {
	var (
		args = os.Args
		l    = len(args) - 1
	)
	// need 4 different cases 1) no args given so we ask again
	// 2) they gave us just one argument so we print there is one
	// 3) they gave us two  arguments so we print there is two
	// 4) they ave us more than two arguments so we print there are %d
	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf(`There are two: "%s %s"`+"\n", args[1], args[2])
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}
