/* panic & recovery */
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something is wrong. contact the administrator")
			return
		} else {
			fmt.Println("Thank you!")
		}
	}()
	result := divide(100, 0)
	fmt.Println("result = ", result)
}

func divide(x, y int) (result int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	result = x / y
	return
}
