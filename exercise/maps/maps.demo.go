package main

import "fmt"

func demo() {
	shoppingList := make(map[string]int)
	// "eggs" are the key and "11" are the value
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	// here we are using the default value which is 0, so 0 += == 0 + 1 = 1
	shoppingList["bread"] += 1
	// Here we are to add one more egg to it, that way we end up with one dozen eggs
	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup:", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "items on the shopping list")

}
