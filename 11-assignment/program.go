package main

import (
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/* implementing the fmt.Stringer interface */
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
Write the following methods

    IndexOf => return the index of the "given product" in the collection (return -1 if the given product is not in the given collection)

    Includes => return true if the given product is present in the collection else return false

    Filter => return products that matches the given criteria
        Use cases:
            1. filter all costly products (cost > 1000)
            2. filter all stationary products (category == stationary)
            3. filter all costly stationary products


    All => return true if ALL the products matches the given criteria else return false
        Use cases:
            1. Are all the products costly products (cost > 1000) ?
            2. Are all the products stationary products? (category == "Stationary")

    Any => return true if ANY of the products matches the given criteria else returns false
        Use cases:
            1. Are there ANY costly product (cost > 1000) ?
            2. Are there ANY stationary products? (category == "Stationary")

	implement sort functionality for the Products
	IMPORTANT : use sort.Sort() function
	Use Cases:
		1. Sort the products by Id
		2. Sort the products by Name
		3. Sort the products by Cost
		4. Sort the products by Units
		5. Sort the products by Category

*/

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
}
