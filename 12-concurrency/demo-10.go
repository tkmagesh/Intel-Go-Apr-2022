package main

import (
	"fmt"
	"sync"
)

//Share memory by communicting

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, &wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println("Result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
