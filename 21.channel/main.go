package main

import "fmt"

func main() {
	fmt.Println("Welcome to channel")

	message := make(chan string)

	go func() {
		message <- "Learn channel"
	}()

	accept := <-message
	fmt.Println(accept)

}
