package main

import "fmt"

func main() {
	exec(f1)
	exec(f2)

	/*
		add(100,200)
		subtract(100,200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func logOperation(x, y int, fn func(int, int)) {
	fmt.Println("invocation started")
	fn(x, y)
	fmt.Println("invocation completed")
}

/*
func logAdd(x, y int) {
	fmt.Println("invocation started")
	add(x, y)
	fmt.Println("invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("invocation started")
	subtract(x, y)
	fmt.Println("invocation completed")
}
*/

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func exec(fn func()) {
	fn()
}
