package main

import "fmt"

func main() {
	fn()

	greet("Magesh")

	greetUser("Magesh", "Kuppan")

	addResult := add(100, 200)
	fmt.Println("Adding 100 and 200, Result = ", addResult)

	//fmt.Println(divide(100, 7))
	/*
		qt, rd := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d & remainder = %d\n", qt, rd)
	*/

	qt, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d \n", qt)

	_, rd := divide(100, 3)
	fmt.Printf("Dividing 100 by 3, remainder = %d \n", rd)

}

func fn() {
	fmt.Println("fn invoked")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
