// An exercise for iota

package main

import "fmt"

// IOTA starts with zero and increases by 1 after each line but there are some caveats as well.
//	First, let’s see a simple example where iota starts with zero increases by 1 after each line
const (
	a = iota
	b
	c
)

func ex1() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
}

// iota expressions – iota allows expressions which can be used to set any value for the constant
const (
	d = iota
	e = iota + 4
	f = iota * 4
)

func ex2() {
	fmt.Println(d) // 0
	fmt.Println(e) // 5
	fmt.Println(f) // 8
}

// iota increment can be skipped using a blank identifier
const (
	r = iota
	_
	t
	y

	// output:
	// 0
	// 2
	// 3
)

type Size int

// lets give numbers for variable wtih iota
const (
	cola = Size(iota)
	chips
	chocolate
	apple
)

// lets try pic chips and apple
func main() {
	var c Size = 1
	c.toString()
	var a Size = 3
	a.toString()
}
func (s Size) toString() {
	switch s {
	case cola:
		fmt.Println("selected cola")
	case chips:
		fmt.Println("selected chips")
	case chocolate:
		fmt.Println("selected chocolate")
	case apple:
		fmt.Println("selected apple")
	default:
		fmt.Println("Invalid entry")
	}
}

// output:
// selected chips
// selected apple
