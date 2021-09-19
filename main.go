package main

import (
	"errors"
	"fmt"
)

/*
AD: Power/Exponent functionality using method recursion. Currently not prepared
	for non-Real numbers (floats and such).
	@param resultpoint: The pointer of the result variable to store modified basenumber
	@param basenumber: Main integer to be modified
	@param exponent: Used to modify resultpointer's value with basenumber
	@param error: 	An error messenger notifier for handling incorrect exponent values
					(exponent must be >= 0; else error)
*/
func mathExponent(resultpointer *int, basenumber int, exponent int) error {
	if exponent > 1 {
		// Initialize resultpointer value...
		if *resultpointer == 0 {
			*resultpointer = 1
		}
		// ...and continue function
		*resultpointer = *resultpointer * basenumber
		mathExponent(resultpointer, basenumber, exponent-1)
	} else if exponent < 1 && exponent >= 0 {
		*resultpointer = 1
	} else if exponent < 0 {
		return errors.New("ERROR: Undefined for negative value")
	}
	return nil
}

/*
AD: Multiplication functionality using method recursion. Currently not prepared
	for non-Real numbers (floats and such).
	@param resultpoint: The pointer to the result variable to store its modification with basenumber
	@param basenumber: The main integer to be modified
	@param iterator: The iterator interger used to modify resultpointer's value with basenumber
	@param error: 	An error messenger notifier for handling future conditions
*/
func mathMultiplication(resultpointer *int, basenumber int, iterator int) error {
	if iterator >= 1 {
		*resultpointer += basenumber // (3,2) => 3 * 2 == 3 + 3
		mathMultiplication(resultpointer, basenumber, iterator-1)
	}
	return nil
}

/*
AD: Addition functionality using method recursion. This is essentially a wrapper of
	mathMultiplcation() where the modifier is saved in resultpointer's value and then
	iterated once with the basenumber. Currently not prepared. for non-Real numbers
	(floats and such).
	@param resultpoint: The pointer to the result variable; replaced with modifier value in this case
	@param basenumber: The main integer to be modified
	@param modifier: The modifier interger used to modify resultpointer's value with basenumber
	@param error: 	An error messenger notifier for handling future conditions
*/
func mathAddition(resultpoint *int, basenumber int, modifier int) error {
	*resultpoint += modifier
	mathMultiplication(resultpoint, basenumber, 1)
	return nil
}

/*
AD: Controller for math functions. Prompts user for input as to what function they need.
	@param userchoice:
*/
func mathController() {
	// Variables
	result := 0
	userchoice := ""
	varone := 0
	vartwo := 0

	// Variables of user's choice
	fmt.Println("Variable One: ")
	fmt.Scanln(&varone)
	fmt.Println("Variable two: ")
	fmt.Scanln(&vartwo)

	// Operation of user's choice
	fmt.Println("Addition (yes/no): ")
	fmt.Scanln(&userchoice)
	if userchoice == "yes" {
		err := mathAddition(&result, varone, vartwo)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Nope!")
	}
}

func main() {
	//err := mathExponent(&result, 3, 4)
	//err := mathMultiplication(&result, 3, 10)
	mathController()
}
