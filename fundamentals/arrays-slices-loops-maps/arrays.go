package main

import (
	"fmt"
)

func main() {
	var myArr [3]int32 = [3]int32{1, 2, 3}
	// myArr2 := [3]int32{1, 2, 3}
	// myArr3 := [...]int32{1, 2, 3}

	fmt.Println(myArr)
	fmt.Println(myArr[1])
	fmt.Println(myArr[1:3])

	// Changing the element at index 1
	myArr[1] = 0
	fmt.Println(myArr)

	// Print out the address of each element. Note in the output that each element occupies 4 bytes of memory.
	fmt.Println(&myArr[0])
	fmt.Println(&myArr[1])
	fmt.Println(&myArr[2])
}
