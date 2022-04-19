package main

/*
import (
	"fmt"
	"modularity-demo/calculator"
)

func main() {
	fmt.Println("App executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OpCount = ", calculator.OpCount)
}
*/

import (
	"fmt"
	calc "modularity-demo/calculator"
)

func main() {
	fmt.Println(calc.Add(100, 200))
}
