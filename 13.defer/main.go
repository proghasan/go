package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")

	// defer run at the end
	defer fmt.Println("Run first") // defer run at the end
	fmt.Println("Run two")
	fmt.Println("Run three")
	fmt.Println("Run four")
	fmt.Println("Run five")

}
