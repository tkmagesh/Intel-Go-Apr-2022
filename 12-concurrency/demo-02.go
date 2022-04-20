package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1)
	go f1()
	f2()
	wg.Wait() //block the execution until the counter becomes 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invocation - started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 invocation - completed")
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
