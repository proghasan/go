package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer")

	// init value
	myNumber := 10
	fmt.Println("My number is", myNumber)

	changeNumber := &myNumber
	fmt.Println("Changed Number with memory address", changeNumber)
	fmt.Println("Changed Number value", *changeNumber)

	*changeNumber = *changeNumber + 5

	fmt.Println("Pointer value change with all reference", myNumber)

}
