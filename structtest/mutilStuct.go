package main

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Infomation Person
	ManagerID  int
}

type SuccessEmployee struct {
	Person
	ManagerID int
}

func main() {
	var employee Employee
	employee.Infomation.FirstName = "John"
	fmt.Println(employee)
	var employee1 SuccessEmployee
	employee1.FirstName = "David"
	fmt.Println(employee1)
}
