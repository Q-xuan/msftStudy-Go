package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")
	highlow(high, low+1)
}

func main() {
	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
