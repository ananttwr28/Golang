/*
----------------- 31/01/26 ----------------

----------------  Lec 7 -- Constants in Golang  ----------------


Constants -- Used when a variable should have a fixed value that cannot be changed, you can use the const keyword. 


*/

package main

import "fmt"

/* 
const and var both can be assigned outside the function, but cannot use the shorthand declaration (:=)
ex - name := "laxman"		// this is not allowed
*/

const age = 15			// and const or var declared outside fn. even not used doesn't give any error
var name = "alex"


func main() {
	/* 
	convention for defining constant (M1)
	const(keyword) const_name const_type  = value
	same other methods are also used same as variable
	*/

	const name string = "ram"			// value has to be assigned to constant at time of declaration only cannot assign later
	// name = "sita"		// this gives an error -- cannot assign to name

	fmt.Println(name)		// prints ram as var inside fn has more priority over global

	fmt.Println(age)


	// constant grouping -- if have to declare multiple constants at the same time.

	const (
		port = 8000
		host = "localhost"
	)

	fmt.Println(port, host)

}
