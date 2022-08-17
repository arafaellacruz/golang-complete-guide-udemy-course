//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.

package main

import (
	"fmt"
)

//* Write a function that returns any two numbers
func twoNumbers() (int, int) {
	return 42, 100
}

//* Write a function that returns any number
func anyNumber() int {
	return 42
}

//* Write a function to add 3 numbers together, supplied as arguments, and return the answer
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

//* Write a function that accepts a person's name as a function parameter and displays a greeting to that person
func helloPerson(name string) {
	fmt.Println("Hello", name)
}

//* Write a function that returns any message, and call it from within fmt.Println()
func hiThere() string {
	return "Here you have any message!"
}

//* Call every function at least once
func main() {
	helloPerson("Rafaella")
	fmt.Println(hiThere())

	//* Add three numbers together using any combination of the existing functions.
	a, b := twoNumbers()
	result := addThreeNumbers(anyNumber(), a, b)

	//* Print the result
	fmt.Println(result)

}
