/*
----------------- 31/01/26 ----------------

----------------  Lec 10 -- switch in Golang  ----------------

switch -- used when multiple if and else if are there
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple Switch Condition
	fmt.Println("Simple Switch Condition")

	var i int
	fmt.Print("Enter the value to be checked: ")
	fmt.Scan(&i)

	switch i {			// switch (keyword) i (val to be checked)
	case 1:				// case (1) the values i can have to be matched
		fmt.Println("value is one")
		// break -- doen't have to write in go, as other lang, it doesn't prints anything below 1 it handles internally
	case 2:
		fmt.Println("value is two")
	case 3:
		fmt.Println("value is three")
	default:
		fmt.Println("other")
	}

	// Multiple Condition Switch
	fmt.Println("Multiple Condition Switch")

	switch time.Now().Weekday() {		// time is std. package in which methods are Now and Weekday
	// time.Now() -- gives current local time based on the system and .Weekday() on it gives which day the current time falls
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	// Type Switch -- A type switch in Go is a specialized form of the switch statement used to determine the concrete type of an interface variable at runtime and execute different blocks of code based on that type. It's a cleaner, more idiomatic way to perform multiple type assertions. 

	fmt.Println("Type Switch")

	// interface{} -- interface empty means the var can hold any type, this is called interface variable 	
	// Fn  in go are first class citizens -- This means they can be treated like any other variable: they can be assigned to variables, passed as arguments to other functions, and returned from functions. 

	whoAmI := func(i interface{}) {			// this is assigning fn to variables
		switch i.(type) {				// declaration of i var is done inside switch and it will take the type of i from i.type
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")	
		default:
			fmt.Println("other data type")	
		}
	}
	whoAmI(50)
}	

