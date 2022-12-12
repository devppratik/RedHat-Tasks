package main

import (
	"fmt"
	"packages/random"
)

// SubTask 1 Main Code. Returns the random number generated
// This is extended in subTask 2 and 3
func taskOne() int64 {
	// Get the random number
	num := random.GetRandom()
	fmt.Printf("Random Number in [1, 100] is : %d\n", num)
	if num > 50 {
		fmt.Print("It's closer to 100 ")
	} else {
		fmt.Print("It's closer to 0")
	}
	return num
}

// Sub Task 2. Gets the num from taskOne()
func taskTwo() {
	num := taskOne()
	if num == 50 {
		fmt.Println("It's 50")
	}
}

// Sub Task 3. Gets the num from taskOne()
func taskThree() {
	num := taskOne()
	if num > 50 && num%2 == 0 {
		fmt.Print("and it's even!")
	}
}
