package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (e Employee) WhoAmI() {
	fmt.Println("I am the Employee ", e.Name)
}

func (emp Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %f", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) ApplyBonus(bonusPercentage float32) {
	emp.Salary = emp.Salary + emp.Salary*(bonusPercentage/100)
}

type FulltimeEmployee struct {
	Employee
	Benefits string
}

//overriding the Employee.Format() method
func (fte FulltimeEmployee) Format() string {
	return fmt.Sprintf("%v, Benefits= %q", fte.Employee.Format(), fte.Benefits)
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
	fmt.Println(fte.Employee.Name)
	fmt.Println(fte.Name)

	fmt.Println(fte.Format())

	fmt.Println("After applying bonus")
	fte.ApplyBonus(20)
	fmt.Println(fte.Format())
}
