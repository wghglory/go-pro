package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// -----------------------------Array-----------------------------------
	// var names [3]string
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArr := names                  // copy value, different address
	names[0] = "Canoe"                 // change names, should not affect otherArr
	fmt.Println("names:", names)       // names: [Canoe Lifejacket Paddle]
	fmt.Println("otherArr:", otherArr) // otherArr: [Kayak Lifejacket Paddle]

	refArr := &names // reference by pointer
	names[0] = "Derek"
	fmt.Println("names:", names)    // names: [Derek Lifejacket Paddle]
	fmt.Println("refArr:", *refArr) // refArr: [Derek Lifejacket Paddle]

	//-------------------Compare Array---------------------
	a := [3]string{"Kayak", "Lifejacket", "Paddle"}
	b := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := a == b
	fmt.Println("comparison:", same) // true!!! array are equal if same type and contain equal elements in the same order. Difference with javascript... // console.log([1]===[1]) will be false

	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}

	// --------------------------------Slices --------------------------------
	// nameSlice := make([]string, 3, 6) // caution about index out of range error
	// nameSlice[0] = "Kayak"
	// nameSlice[1] = "Lifejacket"
	// nameSlice[2] = "Paddle"
	// fmt.Println("len:", len(nameSlice)) // 3
	// fmt.Println("cap:", cap(nameSlice)) // 6

	// --------------------------------append to original--------------------------------
	// nameSlice := []string{"Kayak", "Lifejacket", "Paddle"}
	// nameSlice = append(nameSlice, "Hat", "Gloves")
	// fmt.Println(nameSlice) // [Kayak Lifejacket Paddle Hat Gloves]

	/* ------------------------------Allocating Additional Slice CAPACITY---------------------
	// 2 slices, modify first doesn't affect second
	// nameSlice := []string{"Kayak", "Lifejacket", "Paddle"}
	// newSlice := append(nameSlice, "Hat", "Gloves")
	// nameSlice[0] = "Canoe"
	// fmt.Println("nameSlice:", nameSlice) // nameSlice: [Canoe Lifejacket Paddle]
	// fmt.Println("newSlice:", newSlice)   // newSlice: [Kayak Lifejacket Paddle Hat Gloves]

	// nameSlice := make([]string, 3)
	// nameSlice[0] = "Kayak"
	// nameSlice[1] = "Lifejacket"
	// nameSlice[2] = "Paddle"
	// // capacity is 3, already 3 elements, so when append nameSlice, have to create a new slice. CAPACITY is important
	// newSlice := append(nameSlice, "Hat", "Gloves")
	// nameSlice[0] = "Canoe"
	// fmt.Println("nameSlice:", nameSlice) // [Canoe Lifejacket Paddle]
	// fmt.Println("newSlice:", newSlice)   // [Kayak Lifejacket Paddle Hat Gloves]

	nameSlice := make([]string, 3, 6)
	nameSlice[0] = "Kayak"
	nameSlice[1] = "Lifejacket"
	nameSlice[2] = "Paddle"
	// capacity is 6, already 3 elements, so when append nameSlice, original slice still exists, use the same array (backed by the same underlying array)
	stillThatSlice := append(nameSlice, "Hat", "Gloves")
	nameSlice[0] = "Canoe"
	fmt.Println("nameSlice:", nameSlice)           // [Canoe Lifejacket Paddle]
	fmt.Println("stillThatSlice:", stillThatSlice) // [Canoe Lifejacket Paddle Hat Gloves]
	*/

	// -------------------------------Append one slice to another---------------------------------
	slice1 := make([]string, 3, 6)
	slice1[0] = "Kayak"
	slice1[1] = "Lifejacket"
	slice1[2] = "Paddle"
	slice2 := []string{"Hat Gloves"}
	appendedNames := append(slice1, slice2...)   // ... variadic parameter
	fmt.Println("appendedNames:", appendedNames) // [Kayak Lifejacket Paddle Hat Gloves]

	// -------------------------------create slice from existing array------------------------------------
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// someNames := products[1:3] // [Lifejacket Paddle]
	// allNames := products[:]    // [Kayak Lifejacket Paddle Hat]
	// fmt.Println("someNames:", someNames)
	// fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // someNames len: 2 cap: 3
	// fmt.Println("allNames", allNames)
	// fmt.Println("allNames len", len(allNames), "cap:", cap(allNames)) // allNames len 4 cap: 4

	// allNames := products[1:]              // ["Lifejacket", "Paddle", "Hat"]
	// someNames := allNames[1:3]            // ["Paddle", "Hat"]
	// allNames = append(allNames, "Gloves") // ["Lifejacket", "Paddle", "Hat", "Gloves"]
	// allNames[1] = "Canoe"                 // ["Lifejacket", "Canoe", "Hat", "Gloves"]
	// fmt.Println("someNames:", someNames) // [Paddle Hat]
	// fmt.Println("allNames", allNames)    // [Lifejacket Canoe Hat Gloves]

	// ---------------------- copy Function to Ensure Slice Array Separation-----------------------
	// allNames := products[1:]
	// someNames := make([]string, 2)
	// copy(someNames, allNames)            // copy(dest, source)
	// fmt.Println("someNames:", someNames) // [Lifejacket Paddle]
	// fmt.Println("allNames", allNames)    // [Lifejacket Paddle Hat]
	// // NOTE: must initialize the dest slice before copying! uninitialized slices have 0 length and zero capacity.
	// // copy stops copying when the destination length is reached, since the length is zero, no copying occurs.
	// // var someNames []string won't copy since it's not initialized

	// --------------------------------------sort slice---------------------------------------------
	sort.Strings(products[:])
	for index, value := range products {
		fmt.Println("index", index, "value", value)
	}

	//--------------------------------comparing slice--------------------------------
	// slices can be compared only to the nil value using ==. DeepEqual can compare slices
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))

	// --------------------------------Getting the Array Underlying a Slices --------------------------------
	arrayPtr := (*[3]string)(p1) // array length should be <= slice length
	array := *arrayPtr
	fmt.Println(array)

	// -------------------------------- Maps ----------------------------------
	productMap := make(map[string]float64, 10)
	productMap["Kayak"] = 279
	productMap["Lifejacket"] = 48.8
	productMap["Hat"] = 0

	if value, ok := productMap["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
	delete(productMap, "Hat")

	// --------------------------------Enumerating a Map in Order--------------------------------
	keys := make([]string, 0, len(productMap))
	for key := range productMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key: ", key, "Value: ", productMap[key])
	}
}
