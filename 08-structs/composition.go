package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

type FulltimeEmployee struct {
	Employee
	Benefits string
	Dummy
}

type Dummy struct {
	Name string
}

func main() {
	fte := FulltimeEmployee{
		Employee: Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		},
		Benefits: "Healthcare",
	}
	fmt.Println(fte)
	//fmt.Println(fte.Employee.Name)
	//fmt.Println(fte.Name)
	fmt.Println(fte.Employee.Name)
}
