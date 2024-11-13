package main

import "fmt"

type Notify interface {
	Send(message string) error
}

// For sended to Email
type EmailNotify struct {
	EmailAdress string
}

// For sendend to Number
type PnoneNotify struct {
	PhoneNumber string
}

func (e EmailNotify) Send(message string) error {
	// Some logic send Email message
	fmt.Printf("Send message to Email %s: %s\n", e.EmailAdress, message)
	return nil
}

func (s PnoneNotify) Send(message string) error {
	// Some logic send Phonenumber message
	fmt.Printf("Send message to phone number %s: %s\n", s.PhoneNumber, message)
	return nil
}

func NotifyUser(n Notify, message string) {
	err := n.Send(message)
	if err != nil {
		fmt.Println("Error fail to send message for user")
		return
	}
}

func main() {
	emailNotifer := EmailNotify{EmailAdress: "test@gmail.com"}
	phoneNotifer := PnoneNotify{PhoneNumber: "+372 0000 0000"}

	message := "Your transaction is sucussesful!"

	NotifyUser(emailNotifer, message)
	NotifyUser(phoneNotifer, message)
}
