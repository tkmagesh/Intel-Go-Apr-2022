package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

/*
func (Employee) WhoAmI() {
	fmt.Println("I am an Employee")
}
*/

func (e Employee) WhoAmI() {
	fmt.Println("I am the Employee ", e.Name)
}

func (emp Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %f", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) ApplyBonus(bonusPercentage float32) {
	emp.Salary = emp.Salary + emp.Salary*(bonusPercentage/100)
}

func main() {

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}

	fmt.Printf("%#v\n", emp)
	emp.WhoAmI()
	//fmt.Println(Format(emp))
	fmt.Println(emp.Format())
	fmt.Println("After 10% bonus")
	//ApplyBonus(&emp, 10)
	emp.ApplyBonus(10)
	fmt.Println(emp.Format())
	/*
		e2 := new(Employee)
		fmt.Printf("%T\n", e2)
		fmt.Println(e2.Id) */
}
