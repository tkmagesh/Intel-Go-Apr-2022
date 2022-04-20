package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	goRoutineCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("main started")
	for i := 1; i <= goRoutineCount; i++ {
		wg.Add(1)
		go f1(i)
	}
	f2()
	wg.Wait() //block the execution until the counter becomes 0
	//time.Sleep(5 * time.Second)
	fmt.Println("main completed")
}

func f1(id int) {
	fmt.Printf("f1 [%d] invocation - started\n", id)
	time.Sleep(5 * time.Second)
	fmt.Printf("f1 [%d] invocation - completed\n", id)
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
