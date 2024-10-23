package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("main():")
	log.Print("Hey, I'm a log")

	log.Fatal("Hey, I'm a error log !")
	fmt.Println("Can you see me?")
}
