package main

import "fmt"

type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	rs := Person{"John Doe", "China"}
	zllsb := Person{"Zll", "USA"}
	fmt.Printf("%s\n%s\n", rs, zllsb)
}
