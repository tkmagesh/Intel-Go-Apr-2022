package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var counter int64

func main() {
	goRoutineCount := 10
	fmt.Println("main started")
	for i := 1; i <= goRoutineCount; i++ {
		wg.Add(1)
		go f1(i)
	}
	f2()
	wg.Wait()
	fmt.Println("main completed")
	fmt.Printf("counter = %d\n", counter)
}

func f1(id int) {
	atomic.AddInt64(&counter, 1)
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
