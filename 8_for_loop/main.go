/*
----------------- 31/01/26 ----------------

----------------  Lec 8 -- for loop in Golang  ----------------

for is the only construct in golang for looping there is no while and do while
*/

package main

import "fmt"

func main() {
	// while loop using for

	i := 1		// initialise a var just like other lang.
	fmt.Println("-------- while loop ------------") 
	for i <= 5 {		// condition check without ()
		fmt.Println(i)		// do some work
		i = i + 1			// increment the counter
	}


	// infinite loop 
	fmt.Println("-------- Infinite loop ------------") 
	/* 
	for {
		fmt.Println("go")
	}
	*/

	// for loop
	fmt.Println("-------- for loop ------------") 

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// break and continue keyword

	fmt.Println("-------- break and continue keyword -----------") 

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue		// doesn't print/ skips the curr iteration i.e. 2
		}
		fmt.Println(i)

		if i == 10 {
			break			// breaks from the loop as soon as this condition is met
		}
	}

	// range keyword from go version 1.22+

	fmt.Println("-------- range keyword -----------") 

	for i := range 4 {		// starts from 0 and goes less than the defined value
		fmt.Println(i)
	}

	
}	
