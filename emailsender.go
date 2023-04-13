package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

func main() {
	// string that contains the subject of the email

	// string slice that containes the email adresses
	to := getEmailAdresses()

	// byte slice that containes the the content of the message to be sent
	Subject := []byte("Subject : test\r\n")
	msg := append(Subject,getEmailContent()...)
		//sender Data
	from := "mekchedy7@gmail.com"
	password := "cjohugbcoipxndbj"
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

// reads the file name to read from from the command line and returns a string slice containing the emailadresses
// This function doesn't check for the correctness of the input nor of the emails are correctly written in the file

func getEmailAdresses() []string {
	var filename string
	fmt.Println("Give me the file name to read from the email adresses : ")
	_, err := fmt.Scanln(&filename)
	if err != nil {
		log.Fatal(err)
	}

	eAdrr1, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(eAdrr1), "/n")
}

// reads the file name to read from from the command line and returns a byte slice containing the message to be sent
// This function doesn't check for the correctness of the input

func getEmailContent() []byte {
	var filename string
	fmt.Println("Give me the file name to read from the email content : ")
	_, err := fmt.Scanln(&filename)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}
