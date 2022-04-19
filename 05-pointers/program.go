package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr)
	fmt.Printf("Dereferencing (accessing the value at address %p), no = %d\n", noPtr, *noPtr)

	fmt.Printf("Before increment(), no = %d\n", no)
	increment(&no)
	fmt.Printf("After increment(), no = %d\n", no)

	x, y := 10, 20
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func increment(val *int) {
	*val++
}

func swap(n1, n2 *int) /* DO NOT RETURN any values */ {
	*n1, *n2 = *n2, *n1
}
