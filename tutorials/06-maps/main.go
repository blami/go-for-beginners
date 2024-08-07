package main

import "fmt"

func main() {
	var jpDays map[string]string
	jpDays = make(map[string]string)
	jpDays["日"] = "Sunday"
	jpDays["月"] = "Monday"
	fmt.Println(jpDays)
	fmt.Println(len(jpDays))

	jpRegions := map[string][]string{ // short literal declaration
		"Kanto": {"Tokyo", "Yokohama", "Saitama"},
	}
	fmt.Println(jpRegions, len(jpRegions))

	// checking presence in map
	numbers := map[string]int{
		"one":   1,
		"seven": 7,
		"ten":   10,
	}
	value, ok := numbers["one"]
	fmt.Println(value, ok)
	value, ok = numbers["eight"]
	fmt.Println(value, ok)
}
