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
	swap()
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func increment(val *int) {
	*val++
}

func swap( /*  */ ) /* DO NOT RETURN any values */ {

}
