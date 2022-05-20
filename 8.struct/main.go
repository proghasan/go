package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
	user := User{
		Name:   "Mohammad Hasan",
		Email:  "hasan@gmail.com",
		Status: true,
		Age:    25,
	}

	fmt.Printf("Your name is: %s and email is %s and status %v", user.Name, user.Email, user.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
