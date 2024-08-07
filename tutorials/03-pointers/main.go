package main

import "fmt"

func main() {
	baz := "yes!"

	var myPtr *string
	fmt.Println(myPtr)
	myPtr = &baz
	fmt.Println(myPtr, baz)
	*myPtr = "no!"
	fmt.Println(myPtr, baz)
}
