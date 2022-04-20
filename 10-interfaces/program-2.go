package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

/* implmentation of fmt.Stringer interface */
func (emp Employee) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %f", emp.Id, emp.Name, emp.Salary)
}

func main() {
	e := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Println(e)
}
