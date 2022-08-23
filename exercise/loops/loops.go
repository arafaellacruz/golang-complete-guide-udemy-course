//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.

package main

import "fmt"

func main() {
	//* Print integers 1 to 50
	for i := 1; i <= 50; i++ {

		//* The remainder operator (%) can be used to determine divisibility
		d3 := i%3 == 0
		d5 := i%5 == 0

		//  - Print "Fizz" if the integer is divisible by 3
		//  - Print "Buzz" if the integer is divisible by 5
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		if d3 && d5 {
			fmt.Println("FizzBuzz")
		} else if d3 {
			fmt.Println("Fizz")
		} else if d5 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
