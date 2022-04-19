package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["Pen"] = 2
	*/
	//var productRanks map[string]int = map[string]int{"Pen": 2, "Pencil": 1, "Marker": 3}
	var productRanks map[string]int = map[string]int{
		"Pen":    2,
		"Pencil": 1,
		"Marker": 3,
	}
	productRanks["Scibble-Pad"] = 4
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) = ", len(productRanks))

	fmt.Println()
	fmt.Println("Iterating using range")
	for key, value := range productRanks {
		fmt.Printf("productRanks[%s] = %d\n", key, value)
	}

	fmt.Println()
	fmt.Println("Removing an item (using delete)")
	delete(productRanks, "Marker")
	fmt.Println(productRanks)

	fmt.Println()
	fmt.Println("Checking if a key exists")
	//keyToCheck := "Pen"
	keyToCheck := "Marker"
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Value of %q is %d\n", keyToCheck, val)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}
}
