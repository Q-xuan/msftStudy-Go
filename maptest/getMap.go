package main

import "fmt"

func main() {
	studentAge := make(map[string]int)

	studentAge["john"] = 32
	studentAge["bob"] = 31

	fmt.Println("Bob age is", studentAge["bob"])
	fmt.Println("Christy's age is", studentAge["christy"])

	age, exist := studentAge["christy"]

	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
}
