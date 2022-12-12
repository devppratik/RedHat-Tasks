package main

import (
	"fmt"
	"os"
)

func main() {
	// Swtich case based on the cmd-line arguments of task number
	switch os.Args[1] {
	case "1":
		taskOne()
	case "2":
		taskTwo()
	case "3":
		taskThree()
	// Default case to handle if task number is not in 1 - 3
	default:
		fmt.Println("Wrong Sub-Task Number. Enter between 1 - 3")
	}

}
