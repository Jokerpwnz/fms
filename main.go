package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := modThree("1101")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func modThree(input string) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("input string cannot be empty")
	}

	transitions := [3][2]int{
		{0, 1}, // S0 transitions
		{2, 0}, // S1 transitions
		{1, 2}, // S2 transitions
	}

	currentState := 0 // Start with S0

	for _, char := range input {
		if char != '0' && char != '1' {
			return 0, fmt.Errorf("invalid input: '%c' is not a valid binary digit", char)
		}

		bit := int(char - '0')
		currentState = transitions[currentState][bit]
	}

	return currentState, nil
}