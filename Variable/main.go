package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// one way variable declare
	/*
		var firstNumber int
		var secondNumber int
		firstNumber = 10
		secondNumber = 20
		fmt.Printf("first number %d and second number %d \n", firstNumber, secondNumber)
	*/

	// system time
	rand.Seed(time.Now().UnixNano())

	// other way & guessing game
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	// read data from system
	reader := bufio.NewReader(os.Stdin)

	// display welcome message
	fmt.Printf("Welcome to Guessing game \nGuess the number Game \n")
	fmt.Println("--------------------------------------")
	fmt.Println()

	// instruction
	fmt.Println("Think of a number between 1 to 10", prompt)
	reader.ReadString('\n')

	// multiplication
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("And Multiply result by", secondNumber, prompt)
	reader.ReadString('\n')

	// divide number of original thought number
	fmt.Println("And Divide the result by the number your originally thought of", prompt)
	reader.ReadString('\n')

	// subtraction
	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	answer = (firstNumber * secondNumber) - subtraction
	fmt.Println("Your answer is", answer)
}
