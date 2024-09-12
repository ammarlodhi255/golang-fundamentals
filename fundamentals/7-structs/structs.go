package main

import (
	"fmt"
)

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type owner struct {
	name string
}

// We can add the fields of the owner struct as subfields to another struct
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner // we have a field named owner which is also of type owner.
}

// How to add methods?
func (engine gasEngine) milesLeft() uint8 {
	return engine.mpg * engine.gallons
}

func main() {
	var myEngine gasEngine = gasEngine{8, 21, owner{"Ammar"}}
	var myEngine2 gasEngine = gasEngine{mpg: 10, gallons: 25, ownerInfo: owner{"Ammar"}}
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg)
	fmt.Println(myEngine.gallons)
	fmt.Println(myEngine.ownerInfo.name)

	myEngine2.mpg = 11

	fmt.Println(myEngine2)

	// Anonymous structs:
	var myAnonEngine = struct {
		mpg     uint8
		gallons uint8
	}{2, 4}

	fmt.Println(myAnonEngine)

	// Create an electricEngine instance
	var myElecEngine = electricEngine{mpkwh: 8, kwh: 21, owner: owner{"Ammar"}}

	fmt.Println(myElecEngine)

	// How to call the methods?

	fmt.Println(myEngine.milesLeft())
}
