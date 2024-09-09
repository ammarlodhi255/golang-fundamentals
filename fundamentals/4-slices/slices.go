package main

import (
	"fmt"
)

func main() {
	// Slices wrap arrays. They can be thought of as dynamic arrays.
	myIntSlice := []int32{1, 2}

	fmt.Println(myIntSlice)
	fmt.Println(len(myIntSlice), cap(myIntSlice))

	myIntSlice = append(myIntSlice, 8)

	fmt.Println(myIntSlice)
	fmt.Println(len(myIntSlice), cap(myIntSlice)) // Although a single element was added, the capacity increased by 2: {1, 2, 8, *}

	// You can also use append function to concatenate two slices.
	myIntSlice2 := []int32{1, 2}
	myIntSlice3 := []int32{}
	myIntSlice3 = append(myIntSlice, myIntSlice2...)

	fmt.Println(myIntSlice3)

	// Another way to create slices is through a make function

	myIntSlice4 := make([]int32, 2, 8) // (type, num_of_elems, capacity)
	fmt.Println("myIntSlice4 before:", myIntSlice4)

	myIntSlice4[0] = 1
	myIntSlice4[1] = 2
	fmt.Println("myIntSlice4 after:", myIntSlice4)
}
