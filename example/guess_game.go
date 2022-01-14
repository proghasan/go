package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type number in, just press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstName = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	// reads inputs from standard in the keyboard
	reader := bufio.NewReader(os.Stdin)

	// display welcome/instrations
	fmt.Println("Guess the number Game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of number between 1 to 10", prompt)
	reader.ReadString('\n')

	// take game through the games
	fmt.Println("Multiply your number by", firstName, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number your original thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now Subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	answer = firstName*secondNumber - subtraction
	fmt.Println("The result is", answer)

}
