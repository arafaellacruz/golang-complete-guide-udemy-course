//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.

package main

import "fmt"

type securityTag bool

//* Create a structure to store items and their security tag state
type Items struct {
	name string
	tag  securityTag
}

//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
func activate(tag *securityTag) {
	*tag = true
}

func inactive(tag *securityTag) {
	*tag = false
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(item []Items) {
	for i := 0; i < len(item); i++ {
		inactive(&item[i].tag)
	}
}

func main() {
	//  - Create at least 4 items, all with active security tags
	headphone := Items{"Headphones", true}
	macbook := Items{"Macbook", true}
	mouse := Items{"Mouse", true}
	keyboard := Items{"Keyboard", true}

	//  - Store them in a slice or array
	items := []Items{headphone, macbook, mouse, keyboard}
	fmt.Println("Initial", items)

	//  - Deactivate any one security tag in the array/slice
	inactive(&items[2].tag)
	fmt.Println("Item 2 deactivated.", items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println("Deactivated all the tags.", items)

}
