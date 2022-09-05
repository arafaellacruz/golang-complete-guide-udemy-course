//--Summary:
//  Create a program to manage parts on an assembly line.

//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func showAssemblyLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	//  - Create an assembly line having any three parts
	assembly := make([]Part, 3)
	assembly[0] = "Airbag"
	assembly[1] = "Seat"
	assembly[2] = "Accelerator"

	//* Create a function to print out the contents of the assembly line
	fmt.Println("The first three parts of assembling a car:")
	showAssemblyLine(assembly)

	//* Perform the following:
	//  - Add two new parts to the line
	assembly = append(assembly, "Battery", "Air filter")
	fmt.Println("\nAdded two parts: ")
	showAssemblyLine(assembly)
	//  - Slice the assembly line so it contains only the two new parts
	assembly2 := assembly[3:]
	fmt.Println("\nThe last parts inserted in the assembly line were:")
	showAssemblyLine(assembly2)
}
