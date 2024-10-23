package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	osDefer()
}

func forDefer() {
	for i := 0; i < 4; i++ {
		//压入栈中
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

func osDefer() {
	newFile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}

	defer newFile.Close()

	if _, error = io.WriteString(newFile, "learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file")
		return
	}

	newFile.Sync()
}
