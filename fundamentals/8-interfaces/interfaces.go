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
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) bool {
	return e.milesLeft() >= miles
}

func main() {
	var myEngine gasEngine = gasEngine{8, 21, owner{"Ammar"}}
	var myEngine2 gasEngine = gasEngine{mpg: 10, gallons: 25, ownerInfo: owner{"Ammar"}}
	var myElecEngine = electricEngine{mpkwh: 8, kwh: 21, owner: owner{"Ammar"}}

	var miles uint8 = 90
	fmt.Println(canMakeIt(myEngine, miles))
	fmt.Println(canMakeIt(myEngine2, miles))
	fmt.Println(canMakeIt(myElecEngine, miles))
}
