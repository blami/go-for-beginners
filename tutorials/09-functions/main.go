package main

import "fmt"

func mySum(a, b int) int {
	return a + b
}

func myPositiveSum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("negative a or b")
	}
	return a + b, nil
}

func main() {
	result := mySum(1, 1)
	fmt.Println(result)

	result, err := myPositiveSum(-5, 1)
	if err != nil {
		fmt.Println("error happened!", err)
	}
	fmt.Println(result)
}
