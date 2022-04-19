package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	result, err := divide(100, 7)
	if err == DivideByZeroError {
		fmt.Println("Please donot attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong.", err)
		return
	}
	fmt.Println(result)
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
