package main

import (
	"errors"
	"fmt"
)

// IMPORTS AND VARIEABLES ---

//AD: Power/Exponent functionality using method recursion. Currently not prepared for non-Real numbers (floats and such).
func mathExponent(basenumber int, exponent int) (int, error) {
	result := 1 // Save it in as a class variable
	if exponent > 1 {
		for i := 0; i < exponent; i++ {
			result = result * basenumber
		}
	} else if exponent == 1 {
		result = basenumber
	} else if exponent < 1 && exponent >= 0 {
		result = 1
	} else {
		return -1, errors.New("ERROR: Undefined for negative value")
	}
	return result, nil
}

//AD: Addition functionality using method recursion
func mathAddition(basenumber int, iterator int) int {
	if iterator >= 1 {
		basenumber = mathAddition(basenumber+1, iterator-1)
	}
	return basenumber
}

func main() {
	//fmt.Println("Hello, World!")
	result, err := mathExponent(3, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
