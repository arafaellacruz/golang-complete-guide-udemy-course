//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollTheDices(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {

	// Seed sets the initial value dor the random numbers, UnixNano function will get the number of nanoseconds since the start of the Unix epoch
	rand.Seed(time.Now().UnixNano())

	//  - number of dice used in the rolls
	//  - number of sides of all the dice (6-sided, 10-sided, etc determined
	//    with a variable). All dice must use the same variable for number
	//    of sides, so they all have the same number of sides.
	//  - number of times to roll the dice
	dices, sides := 2, 12
	rolls := 2

	for r := 1; r <= rolls; r++ {
		sum := 0

		for d := 1; d <= dices; d++ {
			rolled := rollTheDices(sides)
			sum += rolled
			fmt.Println("Rolled #", r, ", Die #", d, ": ", rolled)
		}
		fmt.Println("Total rolled: ", sum)

		switch sum := sum; {
		//  - "Snake eyes": when the total roll is 2, and total dice is 2
		case sum == 2 && dices == 2:
			{
				fmt.Println("Snake Eyes!")
			}
		//  - "Even": when the total roll is even
		case sum%2 == 0:
			{
				fmt.Println("Even!")
			}
		//  - "Odd": when the total roll is odd
		case sum%2 != 0:
			{
				fmt.Println("Odd!")
			}
		//  - "Lucky 7": when the total roll is 7
		case sum == 7:
			{
				fmt.Println("Lucky 7!")
			}
		}
	}
}
