package main

import (
	"fmt"
	"sort"
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

//products methods
type Products []Product

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

/*
func (products Products) FilterCostlyProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

/*
func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}
*/

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	}
}

//sort.Interface implementation

func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//Sort by name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//Sort by cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//Sort by units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//Sort by name
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	//IndexOf
	marker := Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println("IndexOf")
	fmt.Println("Index of marker = ", products.IndexOf(marker))

	fmt.Println()
	fmt.Println("Includes")
	fmt.Println("Does the products have a marker ? : ", products.Includes(marker))

	fmt.Println()
	fmt.Println("Filter costly products")
	/*
		costlyProductPredicate := func(p Product) bool {
			return p.Cost > 1000
		}
	*/
	/*
		costlyProductPredicate := getCostlyProductPredicate(1000)
		costlyProducts := products.Filter(costlyProductPredicate)
	*/
	costlyProducts := products.Filter(getCostlyProductPredicate(1000))
	fmt.Println(costlyProducts)

	fmt.Println()
	fmt.Println("Filter stationary products")
	//stationaryProducts := products.Filter(stationaryProductPredicate)
	stationaryProducts := products.Filter(func(product Product) bool {
		return product.Category == "Stationary"
	})
	fmt.Println(stationaryProducts)

	//All
	/* TO DO */

	//Any
	/* TO DO */

	fmt.Println("Default sort (by id)")
	//sort.Sort(products)
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("Sort by name")
	//sort.Sort(ByName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("Sort by cost")
	//sort.Sort(ByCost{products})
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("Sort by units")
	//sort.Sort(ByUnits{products})
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println("Sort by category")
	//sort.Sort(ByCategory{products})
	products.Sort("Category")
	fmt.Println(products)
}

func getCostlyProductPredicate(cost float64) func(p Product) bool {
	return func(p Product) bool {
		return p.Cost > cost
	}
}

func stationaryProductPredicate(p Product) bool {
	return p.Category == "Stationary"
}
