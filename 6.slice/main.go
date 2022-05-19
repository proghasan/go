package main

import "fmt"

func main() {
	fmt.Println("Welcome to slice")

	var fruits = []string{"Apple", "Jack fruit", "Banana"}
	fruits = append(fruits, "New items")
	fmt.Println("Fruits list", fruits)
	fmt.Println()

	// iteration with range
	for index, value := range fruits {
		fmt.Printf("index number %d and value is %s \n", index, value)
	}
}
