//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.

func greeting( a string){
	fmt.Println("Hello ", a)
}


//* Write a function that returns any message, and call it from within
//  fmt.Println()
func anyMessage() string {
	return "This is a return message"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

func addThreeNum(a,b,c int) int {
	return a+b+c
}

//* Write a function that returns any number
func anyNum() int {
	return 1
}

//* Write a function that returns any two numbers
func anyTwoNum() (int, int) {
	return 2,3
}

func main() {

	greeting("Devin")

	fmt.Println(anyMessage())

	a,b := anyTwoNum()
	answer := addThreeNum(a,b,anyNum())

	fmt.Println(answer)

}
