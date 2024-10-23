package main

import "fmt"

func main() {
	q := [...]int{1, 2, 3}
	fmt.Println("arr", q)
	fmt.Println("len", len(q))
	fmt.Println("Capacity", cap(q))
}
