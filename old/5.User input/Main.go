package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var tickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("How much tickets: ")
	fmt.Scan(&tickets)

	fmt.Printf("%v %v thank you buy tickets \n You buy total %v tickets", firstName, lastName, tickets)
}
