package connection

import (
	"fmt"
	"strings"
)

func Pop3Auth(host string, port string, un string, pw string) (*Connection, string) {
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	if check, _ := StringFirstRest(s); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	conn.Write("USER " + un + "\r\n")

	s1, _ := conn.Read()
	if check, _ := StringFirstRest(s1); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	conn.Write("PASS " + pw + "\r\n")

	s2, _ := conn.Read()
	if check, _ := StringFirstRest(s2); check != "+OK" {
		fmt.Println("Login Failed")
		return conn, "err"
	}
	fmt.Println("Login Succeeded")
	return conn, ""
}

func Pop3List(conn *Connection) string {
	conn.Write("list\r\n")
	s, _ := conn.ReadLines(1024)
	first, rest := StringFirstRest(s)
	if first == "+OK" {
		return rest
	}
	return s
}

func StringFirstRest(s string) (string, string) {
	first := strings.Fields(s)[0]
	rest := s[len(first)+1:]
	return first, rest
}
