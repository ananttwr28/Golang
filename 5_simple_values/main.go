/*
----------------- 24/01/26 ----------------

----------------  Lec 5 -- Simple Values in Golang  ----------------


Simple Values -- Basically the data types and values used in Go like int, etc. 

*/

package main

import "fmt"

func main() {

	// Integer
	fmt.Println(1)     // prints 1 it is an int value, println -- print line
	fmt.Println(1 + 1) // can do arithmetic Operations (+,-,*,/) too in print statement

	// String
	fmt.Println("Namma Bengaluru")

	// Bool
	fmt.Println(true)

	// Float
	fmt.Print(10.5, 10, 11) // prints on same line with space in btw by default, but doesn't add a newline at the end.
	fmt.Print(10.6)         // O/p -- 10.5 10 1110.6, No new line after 11 and 10.6

}
