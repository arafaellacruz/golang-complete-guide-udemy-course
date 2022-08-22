//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:

package main

import "fmt"

func main() {
	age := 13

	switch {
	//  - "newborn" when age is 0
	case age == 0:
		{
			fmt.Println("This person are a newborn.")
		}
	//  - "toddler" when age is 1, 2, or 3
	case age >= 1 && age <= 3:
		{
			fmt.Println("This person are a toddler.")
		}
	//  - "child" when age is 4 through 12
	case age >= 4 && age <= 12:
		{
			fmt.Println("This person are a child.")
		}
	//  - "teenager" when age is 13 through 17
	case age == 13 && age <= 17:
		{
			fmt.Println("This person are a teenager.")
		}
	//  - "adult" when age is 18+
	case age >= 18:
		{
			fmt.Println("This person are a adult.")
		}
	}
}
