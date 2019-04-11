package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	//configuration
	hostURL := "smtp.gmail.com"
	hostPORT := "587"
	emailSender := "tiki.marroquin@gmail.com"
	password := "xxxxx"
	emailReciever := "tiki.marroquin@gmail.com"

	//creating auth object
	emailAUTH := smtp.PlainAuth(
		"", emailSender, password, hostURL)

	//creating da email
	msg := []byte("To: " + emailReciever + "\r\n" + "Subject :" + "Hello" + "\r\n" + "How are you doing hoe? im gonna murder you ")

	//send da mail
	err := smtp.SendMail(
		hostURL+":"+hostPORT, emailAUTH, emailSender, []string{emailReciever}, msg)

	if err != nil {
		fmt.Print("Error :", err)
	}
	fmt.Println(" email sent from " + emailReciever)

}
