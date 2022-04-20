package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mutex sync.Mutex
var counter = 0

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
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
