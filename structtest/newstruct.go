package main

import "fmt"

func main() {
	employee := Employee{1001, "john", "Doe", "Doe's Street"}
	fmt.Println(employee)
	employee1 := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee1.LastName)
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}
