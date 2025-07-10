package main

import "fmt"

/**
 * for is Go’s only looping construct. Here are some basic
 * types of for loops.
 **/

func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic “do this N times”
	// iteration is range over an integer.
	for i := range 3 {
		fmt.Println(i)
	}

	// for without a condition will loop repeatedly until you
	// break out of the loop or return from the enclosing function.
	for {
		fmt.Println("looping forever")
		break // remove this line to loop forever
	}

	// You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println(n) // prints odd numbers: 1, 3, 5
	}
}
