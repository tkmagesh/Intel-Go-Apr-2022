package main

import (
	"fmt"
)

//Share memory by communicting

func main() {

	ch := make(chan int)
	fmt.Println("main started")
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
