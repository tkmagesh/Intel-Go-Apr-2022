package main

import "fmt"

func main() {
	//var v interface{}
	var v any
	//v = 100
	//v = "this is a sample string"
	//v = true
	//v = 12.34
	//v = struct{}{}
	v = []int{3, 1, 4, 2, 5}
	fmt.Println(v)

	//type assertion (1)
	/*
		if value, ok := v.(int); ok {
			y := value + 200
			fmt.Println("y = ", y)
		} else {
			fmt.Printf("[%v] is not an int\n", v)
		}
	*/

	//type assertion (2)
	switch value := v.(type) {
	case int:
		y := value + 200
		fmt.Println("v + 200 = ", y)
	case string:
		fmt.Println("len(v) = ", len(value))
	case float64:
		fmt.Printf("[%v] is a float\n", value)
	case bool:
		fmt.Printf("v is a bool, !v = %t\n", !value)
	case struct{}:
		fmt.Println("v is a struct")
	case []int:
		fmt.Println("v is a slice of int")
	default:
		fmt.Println("Unknown type")
	}

}
