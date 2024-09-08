package main

import "fmt"

func main() {
	// int, int16, int32, int64
	var myInt int = 3
	var defaultInt int // this int will default to 32 or 64 bits int depending on your system architecture
	fmt.Printf("Int value: %v\n", myInt)
	fmt.Printf("Default Int value: %v\n", defaultInt)

	// unsigned integer datatypes: uint8, 16, 32, 64
	var myUnsignedInt uint8 = 32
	fmt.Printf("My Unsigned Int Value: %v\n", myUnsignedInt)

	// Note that there is no "float" datatype, you have to define either float32 or float64
	// This is an error: var myFloat float = 1.2
	var myFloat32 float32 = 3.145
	fmt.Printf("myFloat32 value: %v\n", myFloat32)

	// Strings
	var myString string = "Hello World!"
	fmt.Println(myString)

	// Constants
	const PI float32 = 3.145
	fmt.Printf("Value of PI: %v\n", PI)

	// booleans
	var myBool bool = true // or false
	fmt.Println("myBool value:", myBool)

	// Note: GO is strongly typed language.
	// You can't perform operations on two different datatypes without first doing type casting.
	// This is an error: fmt.Println(myInt + myFloat32)

	fmt.Println(float32(myInt) + myFloat32)

	// The len() function returns the number of bytes rather than the number of characters in a string.
	// Since GO uses utf-8 for normal characters, the number of bytes equals the number of characters in that case.

	var normalString string = "Ammar"
	var specialString string = "Gjøvik"

	fmt.Println(len(normalString))
	fmt.Println(len(specialString)) // This outputs 7 but clearly, the length is NOT 7. This is because the number of bytes required for the character ø is 2.
}
