package main

import "fmt"

func main() {
	var myInt1 int // NOTE: zero value
	var myInt2 int = 1
	myInt3 := 2 // short declaration with assignment
	myInt3 = 4  // assignment
	fmt.Println(myInt1, myInt2, myInt3)

	var myBool bool
	var myString string
	fmt.Println(myBool, myString)

	var foo, bar, baz = 42, true, "hi!" // multiple declaration
	fmt.Println(foo, bar, baz)
}
