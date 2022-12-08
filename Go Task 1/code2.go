// Time Complexity O(n)
// Space Complexity O(n)

package main

import "fmt"

func main() {
	// Declare a Slice with few items
	Menu := []string{"Pizza", "Fries", "Coke", "Tacos"}
	// Append to the Slice
	Menu = append(Menu, "Hamburger", "Salad")
	// Loop through the slice
	for _, n := range Menu {
		fmt.Printf("Food: %s\n", n)
	}
}
