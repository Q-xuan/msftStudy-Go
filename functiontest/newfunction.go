package main

import "fmt"

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("perimeter (triangle)", t.perimeter())
	fmt.Println("perimeter (square)", s.perimeter())

	t.doubleSize()
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())
}

type triangle struct {
	size int
}

type square struct {
	size int
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t triangle) perimeter() int {
	return t.size * 3
}

// pointer in function
func (t *triangle) doubleSize() {
	t.size *= 2
}
