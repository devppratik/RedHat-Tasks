// Time Complexity O(n)
// Space Complexity O(1)
package main

import "fmt"

func main() {
	isExit := "" // Exit Case Variable
	for isExit != "exit" {
		var name, nationality string // String Variables to store Name & Natinality
		var age int                  // Integer Variables to store Age
		fmt.Print("Enter Your Name: ")
		fmt.Scan(&name)
		for true { // Loop for integer input validation
			fmt.Print("Enter Your Age: ")
			_, err := fmt.Scanf("%d", &age) // Taking input for age and storing the error
			if err == nil && age > 0 {      // If no error in input i.e an positive integer, break loop
				break
			}
			fmt.Print("Not a valid age - try again\n\n")
			var dump string   // String variable to store the remaining inputs to discard
			fmt.Scanln(&dump) // Store and discard the input buffer
		}
		fmt.Print("Enter Your Nationality: ")
		fmt.Scanln(&nationality)
		fmt.Printf("Your name is %s. Your age is %d years old and your nationality is %s \n", name, age, nationality)
		fmt.Print("\nPress Enter to continue or Type exit to exit ")
		fmt.Scanln(&isExit)

	}
}
