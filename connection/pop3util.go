package connection

import(

 "fmt"
 
 )

func Pop3Auth(host string, port string, un string, pw string) {
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	fmt.Println(s)
	
	conn.Write("USER " + un + "\r\n")
	s1 ,_ := conn.Read()
	fmt.Println(s1)
	conn.Write("PASS " + pw + "\r\n")
	s2 ,_ := conn.Read()
	fmt.Println(s2)
	
	conn.Close()
}
