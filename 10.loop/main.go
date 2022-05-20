package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop")

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// loop
	fmt.Println("For loop starting...")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// range
	fmt.Println("Range without value loop starting...")
	for dayIndex := range days {
		fmt.Println(days[dayIndex])
	}

	// range
	fmt.Println("Range value and index loop starting...")
	for _, day := range days {
		fmt.Printf("Day is %v \n", day)
	}

}
