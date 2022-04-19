package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/cznic/mathutil"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("App executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OpCount = ", calculator.OpCount)
	fmt.Println("Is 21 an even number ? ", utils.IsEven(21))
	color.Yellow("This text will be displayed in yellow color!")
	fmt.Println(mathutil.MaxVal(10, 20, 30, 40))
}

/* import (
	"fmt"
	_ "modularity-demo/calculator"
)

func main() {
	fmt.Println("app executed")
} */
