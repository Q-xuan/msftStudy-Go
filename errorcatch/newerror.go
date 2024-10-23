package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found")

func getInfomation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee, err := apiCallEmployee(1000)
	return employee, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func main() {
	employee, err := getInfomation(1000)
	if errors.Is(err, ErrNotFound) {
		//Something is wrong. Do somthing
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Println(employee)
	}
}
