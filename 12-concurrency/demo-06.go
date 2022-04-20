package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

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
	fmt.Printf("counter = %d\n", counter.count)
}

func f1(id int) {
	counter.Increment()
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
