package main

import "fmt"

func main() {
	func() {
		fmt.Println("anon fn invoked")
	}()

	func(userName string) {
		fmt.Printf("hi %s\n", userName)
	}("Magesh")

	addResult := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("Add result = ", addResult)

	quotient, remainder := func(x, y int) (q, r int) {
		q = x / y
		r = x % y
		return
	}(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", quotient, remainder)
}
