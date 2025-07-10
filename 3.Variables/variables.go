package main

import "fmt"

/**
*	In Go, variables are explicitly declared and used by
*	the compiler to e.g. check type-correctness of function calls.
**/
func main() {
	//var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a) // --> initial

	//You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c) // --> 1 2

	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d) // --> true

	//Variables declared without a corresponding initialization
	//are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e) // --> 0

	//The := syntax is shorthand for declaring and initializing a
	//variable, e.g. for var f string = "apple" in this case. This
	//syntax is only available inside functions
	f := "apple"
	fmt.Println(f) // --> apple
}
