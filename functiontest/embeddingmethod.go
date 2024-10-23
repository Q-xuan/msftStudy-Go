package main

import "fmt"

type triangle struct {
	size int
}

type colorredTriangle struct {
	triangle
	color string
}

func (t triangle) perimeter() int {
	return t.size * 4
}

func (t colorredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := colorredTriangle{triangle{3}, "blue"}
	fmt.Println("Size", t.size)
	fmt.Println("Perimeter (colored)", t.perimeter())
	fmt.Println("Perimeter (normal)", t.triangle.perimeter())
}
