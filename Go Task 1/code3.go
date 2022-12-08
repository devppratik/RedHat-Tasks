// Time Complexity O(n)
// Space Complexity O(n)

package main

import "fmt"

func main() {
	// Declare Fixed Size array of 5 items
	arr := [...]string{"Books", "Bag", "Pencil", "Notebooks", "Geometry"}
	// Loop through the elemnts
	for idx, n := range arr {
		fmt.Printf(" This is %s and its index in the array is %d\n", n, idx)
	}
}
