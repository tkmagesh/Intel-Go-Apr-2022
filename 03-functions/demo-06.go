package main

import "fmt"

func main() {
	fn := getFn()
	fn()

	add := getAdd()
	fmt.Println(add(100, 200))
}

func getFn() func() {
	/*
		var fn func()
		fn = func() {
			fmt.Println("fn invoked")
		}
		return fn
	*/

	return func() {
		fmt.Println("fn invoked")
	}
}

func getAdd() func(int, int) int {
	/*
		add := func(x, y int) int {
			return x + y
		}
		return add
	*/
	return func(x, y int) int {
		return x + y
	}
}
