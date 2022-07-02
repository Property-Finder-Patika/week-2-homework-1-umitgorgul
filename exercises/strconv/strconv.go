package main

import (
	"fmt"
	"strconv"
)

func main() {
	//basic examples
	sayi, _ := strconv.Atoi("-42") //string > int
	yazi := strconv.Itoa(-42)      //int > string
	//string to others
	b, _ := strconv.ParseBool("true")        //string > bool
	f, _ := strconv.ParseFloat("3.1415", 64) //string > float
	i, _ := strconv.ParseInt("-42", 10, 64)  //string > int
	u, _ := strconv.ParseUint("42", 10, 64)  //string > uint
	//others to string
	s1 := strconv.FormatBool(true)                 //bool > string
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64) //float > string
	s3 := strconv.FormatInt(-42, 16)               //int > string
	s4 := strconv.FormatUint(42, 16)               //uint > string
	//print
	fmt.Printf("sayi: %d tip: %T\n", sayi, sayi)
	fmt.Printf("yazi: %s tip: %T\n", yazi, yazi)
	fmt.Printf("b: %t tip: %T\n", b, b)
	fmt.Printf("f: %f tip: %T\n", f, f)
	fmt.Printf("i: %d tip: %T\n", i, i)
	fmt.Printf("u: %d tip: %T\n", u, u)
	fmt.Printf("%T %T %T %T", s1, s2, s3, s4)
}
