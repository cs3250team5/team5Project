package connection

import (
	"fmt"
	"strings"
)

func Pop3Auth(host string, port string, un string, pw string) (*Connection, string) {
	// Checks to see if the username is correct
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	if check, _ := StringFirstRest(s); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	conn.Write("USER " + un + "\r\n")
	// Checks to see if password is correct
	s1, _ := conn.Read()
	if check, _ := StringFirstRest(s1); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	conn.Write("PASS " + pw + "\r\n")
	// Checks to see if username and password is correct
	s2, _ := conn.Read()
	if check, _ := StringFirstRest(s2); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	fmt.Println("Login Succeeded")
	return conn, ""
}

func Pop3List(conn *Connection) string {
	// Lists out all emails with data size
	conn.Write("list\r\n")
	s, _ := conn.ReadLines(1024)
	first, rest := StringFirstRest(s)
	if first == "+OK" {
		return rest
	}
	return s
}

func StringFirstRest(s string) (string, string) {
	// fixes bugs were not all email with data size is present 
	first := strings.Fields(s)[0]
	rest := s[len(first)+1:]
	return first, rest
}

func Pop3Retr(conn *Connection, msgNo int, byteNo int) string {
	// Retreves messages
	message := fmt.Sprintf("RETR %d\r\n", msgNo)
	fmt.Println(message)
	conn.Write(message)
	s, _ := conn.ReadLines(byteNo)
	fmt.Printf("Retrieved %d\n%s\n", msgNo, s)
	return s
}
