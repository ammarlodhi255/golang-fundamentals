package main

import (
	"errors"
	"fmt"
)

func main() {

	var number1 int = 9
	var number2 int = 2

	var result, remainder, err = divideNums(number1, number2)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is: %v\n", result)
	} else {
		fmt.Printf("The result of the division is %v with a remainder %v", result, remainder)
	}
}

func divideNums(numerator int, denominator int) (int, int, error) {
	var err error // nil by default

	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
