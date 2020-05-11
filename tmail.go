package main

import (
	"gopkg.in/gomail.v2"
)


func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", dhfrench@vertisoft.com")
	m.SetHeader("To", "dhf0820@gmail.com", "dhf@dhfrench.net")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Don</b> and <i>Don</i>!")
	m.Attach("/Users/dhf/Documents/XR-7765466.pdf")

	d := gomail.NewDialer("mail.smtp2go.com", 587, "dhfrench@vertisof.com", "sEnyedi57m#1")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
	    panic(err)
	}
}