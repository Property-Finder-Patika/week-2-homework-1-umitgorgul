package main

// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

import (
	"fmt"
	"path"
)

func main() {
	dir, _ := path.Split("secret/deneme.txt")

	fmt.Println(dir)
}
