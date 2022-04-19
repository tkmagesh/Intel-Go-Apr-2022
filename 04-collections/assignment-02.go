package main

import (
	"errors"
	"fmt"
)

type operation func(int, int) int

var operations map[int]operation = map[int]operation{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			break
		}
		n1, n2 := getOperands()
		if result, err := doOperation(n1, n2, userChoice); err != nil {
			fmt.Println("Result = ", result)
		} else {
			fmt.Println("Invalid choice")
			continue
		}
	}
}

func doOperation(n1, n2, userChoice int) (result int, err error) {
	if operation, exists := operations[userChoice]; exists {
		result = operation(n1, n2)
		return
	}
	err = errors.New("operation not supported")
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x + y
}

func divide(x, y int) int {
	return x + y
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the numbers:")
	fmt.Scanf("%d %d\n", &n1, &n2)
	return
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice :")
	fmt.Scanln(&userChoice)
	return userChoice
}
