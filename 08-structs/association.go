package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

type Organization struct {
	Name string
	City string
}

func NewEmployee(id int, name string, salary float32, orgName string, orgCity string) Employee {
	return Employee{
		Id:     id,
		Name:   name,
		Salary: salary,
		Org: &Organization{
			Name: orgName,
			City: orgCity,
		},
	}
}

func main() {
	org := Organization{
		Name: "Intel",
		City: "Boston",
	}
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    &org,
	}
	fmt.Printf("%#v\n", emp)

	org.City = "New York"
	fmt.Printf("%#v\n", emp.Org.City)

	e2 := NewEmployee(200, "Gerald", 1000000, "Intel", "Seatle")
	fmt.Println(e2)
}
