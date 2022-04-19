package main

import "fmt"

func main() {
	var nos []int = []int{3, 1, 4, 2, 5}
	fmt.Printf("nos = %v, len(nos) = %d, cap(nos) = %d\n", nos, len(nos), cap(nos))

	var count int
	fmt.Println("Enter the number of values")
	fmt.Scanln(&count)
	var data []int
	data = make([]int, 0, count) //proactively allocate memory to avoid adhoc allocation
	for idx := 0; idx < count; idx++ {
		var input int
		fmt.Println("Enter the value")
		fmt.Scanln(&input)
		data = append(data, input)
		fmt.Printf("data = %v, len(data) = %d, cap(data) = %d\n", data, len(data), cap(data))
	}

	fmt.Println(data)

}
