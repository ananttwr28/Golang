/*
----------------- 31/01/26 ----------------

----------------  Lec 9 -- if else and input in Golang  ----------------
*/

package main

import "fmt"

func main() {
	/*
	fmt.Scanln(&first) -- is used to take input from user in the next line. 
	While fmt.Scan(&first) -- is used to take input from user in the same line. 
	Ampersand -- is necessary to give the reference as to in which variable we have to store the variable.
	basically fmt package has methods that scans and (&) is used to tell the method to where to save the scanned value so memory address of the var is accessed using (&var_name) and it saves there
	*/

	var age int 
	fmt.Print("Enter your age: ")
	fmt.Scan(&age) 
	
	if age >= 18 {
		fmt.Println("You are eligible for DL")
	} else if age >= 13 && age < 18 {				// Logical operators can be also used inside if else (&&, ||, !)
		fmt.Println("You are in your teenage")
	} else {
		fmt.Println("You are a kid right now!")
	}

	// we can also declare a var inside the if construct

	if height := 175; height > 170 {
		fmt.Println("You are selected in ARMY", height)
	}
	// till go 1.22 it doesn't have ternary operator (int index = val > 0 ? val : -val) so have to use if else only 
}	

