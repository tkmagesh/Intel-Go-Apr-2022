/* Higher order functions - functions can be assigned as values to variables */
package main

import "fmt"

func main() {
	/*
		var fn = func() {
			fmt.Println("fn invoked")
		}
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Println("Hi", userName)
	}
	greet("Magesh")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	add = 100
	fmt.Println(add(100, 200))
}
