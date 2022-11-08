package main

import (
	"fmt"
)

type User struct {
	Name     string
	Email    string
	Notifier UserNotifier
}

type UserNotifier interface {
	SendMessage(user *User, message string) error
}

type EmailNotifier struct{}

func (notifier EmailNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Send email to %s with content\n%s\n", user.Email, message)
	return err
}

type SmsNotifier struct{}

func (notifier SmsNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Send sms to %s with content\n%s\n", user.Name, message)
	return err
}

func (user *User) notify(message string) error {
	return user.Notifier.SendMessage(user, message)
}

func main() {
	fmt.Println("Learning interface")
	fmt.Println("==============================")

	user1 := &User{"Hasan", "test@example.com", EmailNotifier{}}
	user2 := &User{"Abdul Rahim", "test@example.com", SmsNotifier{}}

	user1.notify("User verified email send!!!")
	user2.notify("User verified sms send!!!")
}
