package main

import (
	"fmt"
	"time"
)

func MyFunc(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	MyFunc("direct")
	go MyFunc("goroutine")
	go MyFunc("second goroutine")

	time.Sleep(time.Second)
	fmt.Println("done")
}
