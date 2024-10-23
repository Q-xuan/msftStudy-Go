package main

import (
	"fmt"
	"geometrytest/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())

	// fmt.Println("Size", t.size) program will dead.becuace t.size is not public

}
