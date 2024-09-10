package main

import (
	"fmt"
	"strings"
)

func main() {

	myString := "Hello Universe!"
	// Pretty inefficient way to copy a string. At each iteration you are creating a new string.

	var copyString string

	for i := range myString {
		copyString += string(myString[i])
	}

	fmt.Println(copyString)

	// A better way is to use a string builder

	var strBuilder strings.Builder

	for i := range myString {
		strBuilder.WriteString(string(myString[i]))
	}

	var copyStr string = strBuilder.String()

	fmt.Println(copyStr)
}
