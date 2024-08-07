package main

import "fmt"

func main() {
	var numbers2 []int
	fmt.Println(numbers2)
	numbers2 = make([]int, 4) // len, capacity
	fmt.Println(numbers2, len(numbers2), cap(numbers2))
	numbers2[0] = 4
	numbers2[1] = 5

	fmt.Println(numbers2)
	numbers2 = append(numbers2, 6, 7)

	fmt.Println(numbers2)
	fmt.Println(numbers2[2:])
	fmt.Println(numbers2[:2])
	fmt.Println(numbers2[0:1])

	// This shows array backing slice and slices sharing backing array
	a := [3]int{1, 2, 3}
	s1 := a[1:]
	s2 := a[2:]
	fmt.Println(a, s1, s2)
	s2[0] = 7
	fmt.Println(a, s1, s2)
	fmt.Println(len(a))
}
