package main

import "fmt"

func main() {
	fmt.Println("Welcome to method")
	user := User{
		Name:   "Mohammad Hasan",
		Email:  "hasan@gmail.com",
		Status: true,
		Age:    25,
	}
	fmt.Println(user)
	fmt.Println()

	fmt.Println("User name is: ", user.Name)
	user.GetStatus()
	user.SetStatus()
	fmt.Println(user)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

func (u *User) SetStatus() {
	u.Status = false
}
