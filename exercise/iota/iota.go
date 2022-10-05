//--Summary:
//  Create a calculator that can perform basic mathematical operations.

package main

import "fmt"

//* Mathematical operations must be defined as constants using iota
const (
	add = iota
	sub
	mul
	div
)

type operation int

//* Write a receiver function that performs the mathematical operation on two operands
func (o operation) mathematicalOperation(a, b int) int {
	switch o {
	case add:
		return a + b
	case sub:
		return a - b
	case mul:
		return a * b
	case div:
		return a / b
	}
	panic("invalid operation.")
}

//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
func main() {

	fmt.Println("===== O resultado das operações é: =====")

	add := operation(add)
	fmt.Println("A soma dos números é igual a: ")
	fmt.Println(add.mathematicalOperation(2, 2))

	sub := operation(sub)
	fmt.Println("A diferença entre os números é: ")
	fmt.Println(sub.mathematicalOperation(10, 3))

	mul := operation(mul)
	fmt.Println("O produto dos números é: ")
	fmt.Println(mul.mathematicalOperation(3, 3))

	div := operation(div)
	fmt.Println("A divisão dos números é: ")
	fmt.Println(div.mathematicalOperation(100, 2))
}
