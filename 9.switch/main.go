package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch and case")
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("You have move 2 sport")
		fallthrough
	case 3:
		fmt.Println("You have move 3 sport")
		fallthrough
	case 4:
		fmt.Println("You have move 4 sport")
	case 5:
		fmt.Println("You have move 5 sport")
	case 6:
		fmt.Println("You have move 6 sport")
	default:
		fmt.Println("What is that")

	}
}
