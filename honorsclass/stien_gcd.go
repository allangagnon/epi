package main

import (
	"fmt"
)

//Stein Algorithm for Greatest Common Denominator
func Stein(value1, value2 uint) uint {

	if value1 == 0 {return value2}
	if value2 == 0 {return value1}
	if value1 == value2 {return value1}

	var oneu uint = 1
	var value1IsEven = false
	value1IsEven = (value1 & oneu) == 0

	var value2IsEven = false
	value2IsEven = (value2 & oneu) == 0

	//var res uint

	if value1IsEven && value2IsEven { 
		return Stein(value1 >> 1, value2 >> 1) << 1;
	} else if value1IsEven && !value2IsEven { 
		return Stein(value1 >> 1, value2) 
	} else if value2IsEven {
		return Stein(value1, value2 >> 1) 
	} else if (value1 > value2) {
		return Stein((value1 - value2) >> 1, value2)
	} else {
	return Stein(value1, (value2 - value1) >> 1)
	}

}

func main() {

	var res uint
	res = Stein(116150,232704)
	fmt.Println(res)

	res = Stein(230490,498300)
	fmt.Println(res)
}