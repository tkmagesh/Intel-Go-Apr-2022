/* Higher order functions - functions returned as return values */
package main

import "fmt"

func main() {
	/*
		add(100,200)
		subtract(100,200)
	*/
	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	logAdd := logOperation(add)
	logAdd(100, 200)

	logSubtract := logOperation(subtract)
	logSubtract(100, 200)
}

/*
func logOperation(fn func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		fn(x, y)
		fmt.Println("invocation completed")
	}
}
*/

type operationFn func(int, int)

func logOperation(fn operationFn) operationFn {
	return func(x, y int) {
		fmt.Println("invocation started")
		fn(x, y)
		fmt.Println("invocation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
