//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.

package main

import (
	"fmt"
)

type Product struct {
	name  string
	price int
}

func printStats(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += int(item.price)

		// se item.name for diferente de uma string vazia, contar mais um item na variavel 'totalItems'
		if item.name != "" {
			totalItems += 1
		}
	}

	//* Print to the terminal:
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	fmt.Println("Last item on the list: ", list[totalItems-1])
	fmt.Println("Total items: ", totalItems)
	fmt.Println("Total cost: ", cost)
}

func main() {
	//* Using an array, create a shopping list with enough room for 4 products
	//  - Products must include the price and the name
	//* Insert 3 products into the array
	shoppingList := [4]Product{
		{"Meat - 2kg", 20},
		{"Rice - 5kg", 15},
		{"Beans - 2kg", 8},
	}

	printStats(shoppingList)

	//* Add a fourth product to the list and print out the information again
	shoppingList[3] = Product{"Pasta - 6kg", 12}

	printStats(shoppingList)
}
