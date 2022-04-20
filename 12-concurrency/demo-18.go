package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes()
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
}

func genPrimes() chan int {
	ch := make(chan int)
	timeoutCh := time.After(10 * time.Second)
	go func() {
		no := 3

	LOOP:
		for {
			if isPrime(no) {
				select {
				case <-timeoutCh:
					break LOOP
				case ch <- no:
					time.Sleep(500 * time.Millisecond)
				}
			}
			no++
		}
		close(ch)
	}()
	return ch
}

/* func After(duration time.Duration) chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(duration)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
