package connection

import (
	"Team5Project/util"
	"fmt"
	"log"
)

func Pop3Auth(host string, port string, un string, pw string) (*Connection, string) {
	// Checks to see if the username is correct
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	if check, _ := util.FirstRest(s); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	conn.Write("USER " + un + "\r\n")
	// Checks to see if password is correct
	s1, _ := conn.Read()
	if check, _ := util.FirstRest(s1); check != "+OK" {
		log.Fatal("Login Failed - Username Rejected")
		return conn, "err"
	}
	conn.Write("PASS " + pw + "\r\n")
	// Checks to see if username and password is correct
	s2, _ := conn.Read()
	if check, _ := util.FirstRest(s2); check != "+OK" {
		fmt.Println("Login Failed - Password Rejected")
		return conn, "Bad Password"
	}
	fmt.Println("Login Succeeded")
	return conn, ""
}

func Pop3List(conn Connection) string {
	// Lists out all emails with data size
	conn.Write("list\r\n")
	s, _ := conn.ReadLines(1024)
	first, rest := util.FirstRest(s)
	if first == "+OK" {
		return rest
	}
	return s
}

func Pop3Retr(conn Connection, msgNo int, byteNo int) string {
	// Retreves messages
	message := fmt.Sprintf("RETR %d\r\n", msgNo)
	conn.Write(message)
	s, _ := conn.ReadLines(byteNo)
	return s
}

func RetrieveAll(conn Connection, listMap map[int]int) map[int]MailObject {
	// Emails are put into maps based on data size
	finalMap := make(map[int]MailObject)
	for key, value := range listMap {
		s := Pop3Retr(conn, key, value)
		mail := MailFilter(s)
		mail.Num = key
		finalMap[key] = mail
	}
	return finalMap
}

func Pop3Del(conn Connection, I int) {
	message := fmt.Sprintf("DELE %d\r\n", I)
	conn.Write(message)
}
