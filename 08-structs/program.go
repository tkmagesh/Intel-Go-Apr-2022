package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{100, "Magesh", 10000}
	//var emp Employee = Employee{Id: 100, Name: "Magesh"}
	/*
		var emp Employee = Employee{
			Id:   100,
			Name: "Magesh",
		}
	*/
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Printf("%#v\n", emp)
	fmt.Println(Format(emp))
	fmt.Println("After 10% bonus")
	ApplyBonus(&emp, 10)
	fmt.Println(Format(emp))

	e2 := new(Employee)
	fmt.Printf("%T\n", e2)
	fmt.Println(e2.Id)
}

func Format(emp Employee) string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %f", emp.Id, emp.Name, emp.Salary)
}

func ApplyBonus(emp *Employee, bonusPercentage float32) {

	//(*emp).Salary = (*emp).Salary + (*emp).Salary*(bonusPercentage/100)
	emp.Salary = emp.Salary + emp.Salary*(bonusPercentage/100)
}
