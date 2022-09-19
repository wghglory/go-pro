package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "true"
	val2 := "false"
	val3 := "not true"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)

	fmt.Println("bool1:", bool1, b1err) // bool1: false <nil>
	fmt.Println("bool2:", bool2, b2err) // bool2: false <nil>
	fmt.Println("bool3:", bool3, b3err) // bool3: false strconv.ParseBool: parsing "not true": invalid syntax

	if b1err == nil {
		fmt.Println("Parsed value:", bool1) // execute this line
	} else {
		fmt.Println("cannot parse: ", val1)
	}

	val4 := "100"
	// if statements can define an initialization statement!
	if int4, err := strconv.Atoi(val4); err == nil {
		fmt.Println("Parsed value:", int4) // execute this line
		fmt.Println("Int:", strconv.Itoa(int4))
	} else {
		fmt.Println("cannot parse: ", val4)
	}

	val5 := true
	str5 := strconv.FormatBool(val5)
	fmt.Println(str5)
}
