package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// welcome message
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate your pizza between 1 to 10")

	// read from system
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// print result
	fmt.Println("Thanks your rating,", input)

	// convert number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		// plus 1 with result
		fmt.Println("We added 1 rating in your rating, ", numRating+1)
	}
}
