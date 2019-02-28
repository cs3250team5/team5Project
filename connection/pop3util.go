package connection

import (
	"fmt"
	"strings"
)

func Pop3Auth(host string, port string, un string, pw string) {
	conn := MakeConnection(host, port)
	defer conn.Close()
	s, _ := conn.Read()
	if check, _ := StringFirstRest(s); check != "+OK" {
		fmt.Println("Login Failed")
		return
	}
	conn.Write("USER " + un + "\r\n")

	s1, _ := conn.Read()
	if check, _ := StringFirstRest(s1); check != "+OK" {
		fmt.Println("Login Failed")
		return
	}
	conn.Write("PASS " + pw + "\r\n")

	s2, _ := conn.Read()
	if check, _ := StringFirstRest(s2); check != "+OK" {
		fmt.Println("Login Failed")
		return
	}
	fmt.Println("Login Succeeded")
}

func StringFirstRest(s string) (string, string) {
	first := strings.Fields(s)[0]
	rest := s[len(first)+1:]
	return first, rest
}
