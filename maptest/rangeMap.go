package main

import "fmt"

func main() {
	studentAge := make(map[string]int)
	studentAge["john"] = 32
	studentAge["bob"] = 31
	for name, age := range studentAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
