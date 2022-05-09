package main

import "fmt"

func main() {
	var name string
	var tickets uint16
	var bookings []string
	var bookingArray [50]string

	fmt.Printf("Enter your name ")
	fmt.Scan(&name)

	fmt.Printf("How much ticket you want to buy: ")
	fmt.Scan(&tickets)
	bookings = append(bookings, name)
	bookingArray[0] = name

	//array
	fmt.Printf("Array %v \n", bookingArray)

	// slices
	fmt.Printf("Slices %v \n", bookings)
	fmt.Printf("Welcome %v,\nYou buy total ticket %v", name, tickets)
}
