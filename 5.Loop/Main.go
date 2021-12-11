package main

import "fmt"

func main() {
	// single condition loop / like while loop
	i := 0
	for i < 3 {
		fmt.Printf("Loop number is %v \n", i)
		i++
	}

	// for loop
	for j := 0; j < 10; j++ {
		fmt.Printf("For Loop index number %v \n", j)
	}

	// loop break
	for {
		fmt.Println("Loop")
		break
	}
}
