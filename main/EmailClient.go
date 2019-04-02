package main

import (
	"Team5Project/connection"
	"Team5Project/userInterface"
	"fmt"
)

func main() {
	
	//Gathers username and password and call server
	//un, pw := userInterface.GetUsernameAndPassword()
	un, pw := "fbuker", "su$lpHUr.8097"
	conn, err := connection.Pop3Auth("pop.nyx.net", "110", un, pw)
	defer conn.Close()
	if err != "err" {
		s := connection.Pop3List(conn)
		fmt.Println("Messages in Inbox:", s)
	}
	//test the emails
	//conn.Write("RETR 1\r\n")
	//message1, _ := conn.ReadN(3153)
	//fmt.Println(message1)
	n := userInterface.GetMessageNumber()
	fmt.Println("Retrieving Message", n)
	conn.Write("RETR 3\r\n")
	message3, _ := conn.ReadLines(11542)
	mail := connection.InterpretLines(message3)
	fmt.Println("From: ", mail.From)
	fmt.Println("Subject: ", mail.Subject)
	connection.SaveN(mail, 3)
}
