// Time Complexity O(n)
// Space Complexity O(1)
package main

import "fmt"

func main() {
	// Exit Case Variable
	isExit := ""
	for isExit != "exit" {
		// String Variables to store Name & Natinality
		var name, nationality string
		// Integer Variables to store Age
		var age int
		fmt.Print("Enter Your Name: ")
		fmt.Scan(&name)
		// Loop for integer input validation
		for true {
			fmt.Print("Enter Your Age: ")
			// Taking input for age and storing the error
			_, err := fmt.Scanf("%d", &age)
			// If no error in input i.e an positive integer, break loop
			if err == nil && age > 0 {
				break
			}
			fmt.Print("Not a valid age - try again\n\n")
			// String variable to store the remaining inputs to discard
			var dump string
			// Store and discard the input buffer
			fmt.Scanln(&dump)
		}
		fmt.Print("Enter Your Nationality: ")
		fmt.Scanln(&nationality)
		fmt.Printf("Your name is %s. Your age is %d years old and your nationality is %s \n", name, age, nationality)
		fmt.Print("\nPress Enter to continue or Type exit to exit ")
		fmt.Scanln(&isExit)
	}
}
