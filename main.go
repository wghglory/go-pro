package main

import "fmt"

func main() {
	PrintHello()

	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

// revive:disable:exported
func PrintHello() {
	fmt.Println("Hello, Go")
}

// revive:enable:exported
// PrintNumber writes a number using the fmt.Println function
func PrintNumber(number int) {
	fmt.Println(number)
}
