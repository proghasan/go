package main

import "fmt"

func main() {
	// one way
	var name string
	name = "Mohammad Hasan"

	// show only value
	fmt.Println(name)

	// show with value and type
	fmt.Printf("Welcome %v , Your data type is %T \n", name, name)

	// two way
	age := 10
	fmt.Println(age)
}
