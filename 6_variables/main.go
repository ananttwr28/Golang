/*
----------------- 30/01/26 ----------------

----------------  Lec 6 -- Variables in Golang  ----------------


Variables -- Variables are containers for storing data values. 

Go style prefers camelCase for multi-word names (e.g., firstName, totalPrice) over underscores/ snake case (first_name).

*/

package main

import "fmt"

func main() {
	// convention for defining variable (M1)
	// var(keyword) variable_name variable_type  = value

	var name string = "ram"			// this method of declaring a variable is useful when have to define only the var and the value will be assigned at a later stage, but in that case defining the type is mandatory

	var age int
	age = 15
	fmt.Println(age)

	/*

	In go we have to use the variable or the package if defining or importing else it will give below error
	-- this is enforced in go to maintain efficiency and reduce bugs and less memory consumption of unused ones
	
	.\main.go:13:8: "fmt" imported and not used
	.\main.go:19:6: declared and not used: name

	*/
	
	fmt.Println(name)


	// (M2) of declaring a variable is without writing the type golang will automatically infer it. (it will figure out on its own by the given value)

	var isAdult = true
	fmt.Println(isAdult)


	// (M3) of declaring a var when we want to assign some value at that time only, this is shorthand
	babyAge := 12			// var_name := value
	fmt.Println(babyAge)


	// Multi var declaration

	var height, weight = 170, 56
	fmt.Println(height, weight)

	//  or without var keyword

	sex, relationship := "male", "child"
	fmt.Println(sex, relationship)


}
