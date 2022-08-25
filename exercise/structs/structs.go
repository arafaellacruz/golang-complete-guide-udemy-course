//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.

package main

import "fmt"

type Coordinates struct {
	x, y int
}

//* Create a rectangle structure containing its coordinates
type rectangle struct {
	a Coordinates // top left
	b Coordinates // bottom right
}

func width(rect rectangle) int {
	return (rect.b.x - rect.a.x)
}

func length(rect rectangle) int {
	return (rect.a.y - rect.b.y)
}

//* Using functions, calculate the area and perimeter of a rectangle,
func area(rect rectangle) int {
	return width(rect) * length(rect)
}

func perimeter(rect rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

//  - Print the results to the terminal
func printInfo(rect rectangle) {
	fmt.Println("The Area of the rectangle is equals ", area(rect))
	fmt.Println("The Perimeter of the rectangle is equals ", perimeter(rect))
}

func main() {

	//  - The functions must use the rectangle structure as the function parameter
	rect := rectangle{a: Coordinates{4, 8}, b: Coordinates{6, 12}}
	printInfo(rect)

	//* After performing the above requirements, double the size of the existing rectangle and repeat the calculations
	rect.a.y *= 2
	rect.b.x *= 2
	//  - Print the new results to the terminal
	printInfo(rect)

}
