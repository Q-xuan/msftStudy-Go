package main

import "fmt"

func main() {
	studentAge := make(map[string]int)

	studentAge["john"] = 32
	studentAge["bob"] = 31
	delete(studentAge, "john")
	fmt.Println(studentAge)
}
