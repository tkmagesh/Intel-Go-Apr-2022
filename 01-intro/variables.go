package main

import "fmt"

/* var x = "something" */

func main() {

	/*
		var msg string
		msg = "Hello World!"
		fmt.Println(msg)
	*/

	/*
		var msg string = "Hello World!"
		fmt.Println(msg)
	*/

	/*
		//type inference
		var msg = "Hello World!"
		fmt.Println(msg)
	*/
	msg := "Hello World!" /* := syntax can be used only in a function (not for package level variables) */
	fmt.Println(msg)

	//handling more than one variable declarations
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Result of addition="
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Result of addition="
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Result of addition="
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Result of addition ="
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x = 100
		var y = 200
		var str = "Result of addition ="
		var result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Result of addtion = "
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Result of addtion = "
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Result of addition = "
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Result of addition = "
	result := x + y
	fmt.Println(str, result)
}
