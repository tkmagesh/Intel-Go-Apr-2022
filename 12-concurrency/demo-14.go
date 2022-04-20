package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes()
	for {
		if primeNo, ok := <-ch; ok {
			fmt.Println(primeNo)
		} else {
			break
		}
	}
}

func genPrimes() chan int {
	ch := make(chan int)
	go func() {
		no := 3
		count := 10
		for count > 0 {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
				count--
			}
			no++
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
