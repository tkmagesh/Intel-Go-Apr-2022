package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := func() chan int {
		c1 := make(chan int)
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("goroutine-1 completed")
			c1 <- 100
		}()
		return c1
	}()

	ch2 := func() chan int {
		c2 := make(chan int)
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("goroutine-2 completed")
			c2 <- 200
		}()
		return c2
	}()

	for i := 0; i < 2; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		}
	}

}
