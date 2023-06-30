package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

func main() {

	// string slice that containes the email adresses
	to := getEmailAdresses()

	// byte slice that containes the the content of the message to be sent
	Subject := []byte("Subject : Proposition de mise à jour du site web et discussion de nos solutions pour votre entreprise \r\n")
	msg := append(Subject, getEmailContent()...)
	//sender Data
	from := flag.String("from", "", "email adresse to be used to send the email")
	password := flag.String("pwd", "", "password")
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", *from, *password, host)
	err := smtp.SendMail("smtp.gmail.com:587", auth, *from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

// reads the file name to read from from the command line and returns a string slice containing the emailadresses
// This function doesn't check for the correctness of the input nor of the emails are correctly written in the file

func getEmailAdresses() []string {
	var filename *string
	filename = flag.String("ead", "test.txt", "path to all the email text file to whom sent")

	eAdrr1, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(eAdrr1), ";")
}

// reads the file name to read from from the command line and returns a byte slice containing the message to be sent
// This function doesn't check for the correctness of the input

func getEmailContent() []byte {
	var filename *string
	filename = flag.String("ec", "email_approach_firm_french.txt", "path to email text file to be sent")
	msg, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}
