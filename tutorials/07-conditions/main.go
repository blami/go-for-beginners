package main

import (
	"fmt"
	"runtime"
)

func main() {
	// If
	myBool := true
	if myBool {
		fmt.Println("condition met")
	}
	if !myBool {
		fmt.Println("condition met")
	} else {
		fmt.Println("condition not met")
	}

	// conditions always have to evaluate to true/false
	a := 5
	b := 7
	if a >= 3 {
		fmt.Println("a is greater or equal than 3")
	}
	if a >= 3 && b < 8 {
		fmt.Println("ok")
	}
	if a >= 7 || b < 8 {
		fmt.Println("ok")
	}

	// assign and check
	numbers := map[string]int{
		"one":   1,
		"seven": 7,
	}
	// also explain _
	if _, ok := numbers["eight"]; !ok {
		fmt.Println("eight is not in map")
	}

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("something else")
	}
}
