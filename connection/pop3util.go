package connection

import "fmt"

func Pop3Auth(host string, port string, un string, pw string) {
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	fmt.Println(s)
	conn.Close()
}
