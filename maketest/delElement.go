package main

import (
	"fmt"
)

func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}

	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	fmt.Println(slice2)
}
