package main

import (
	"fmt"
)

func main() {
	var myMap map[string]uint8 = map[string]uint8{"Ammar": 25, "Ahmed": 26, "Muzamil": 26}
	fmt.Println(myMap)

	// Another way of creating maps:
	// var myMap2 map[string]uint8 = make(map[string]uint8) // Empty by default

	for key, value := range myMap {
		fmt.Printf("%v: %v\n", key, value)
	}

	// Accessing values through keys:
	fmt.Println(myMap["Ammar"])
	fmt.Println(myMap["Adam"]) // Accessing value that doesn't exist. How do we know if the value is 0 or the element doesn't exist?

	// We can check that by realizing that accessing a value returns two values:

	key := "Adam"
	var age, ok = myMap[key] // ok = false

	if ok {
		fmt.Println(key+"s age is:", age)
	} else {
		fmt.Println("There is no key named", key)
	}

	// Add a key,value pair

	myMap["Adam"] = 24

	// Delete a value using "delete"

	delete(myMap, "Ammar")
	fmt.Println(myMap)
}
