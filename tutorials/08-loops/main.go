package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "tomato"}
	for i := range fruits {
		fmt.Println(i)
	}
	for i, value := range fruits {
		fmt.Println(i, value)
		// note that
		value = "mango"
	}
	fmt.Println(fruits)
	for _, value := range fruits {
		fmt.Println(value)
	}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	jpDays := map[string]string{
		"日": "Sunday",
		"月": "Monday",
	}
	for jp, en := range jpDays {
		fmt.Println(jp, en)
		// note that
		en = "test"
	}
	fmt.Println(jpDays)
}
