package main

import (
	"flag"
	"log"
	"net/smtp"
)

var (
	user = flag.String("user", "", "user for smtp server")
	pass = flag.String("pass", "", "password for smtp server")
)

func main() {
	flag.Parse()
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", *user, *pass, "smtp.gmail.com"),
		*user,
		[]string{"rif@mailinator.com"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
