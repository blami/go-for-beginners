package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)

	strings := [2]string{"hello", "world"}
	fmt.Println(strings)
}
