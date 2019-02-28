package connection

import(

 "fmt"
 
 )

func Pop3Auth(host string, port string, un string, pw string)*Connection {
	
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	fmt.Println(s)
	
	conn.Write("USER " + un + "\r\n")
	s1 ,_ := conn.Read()
	fmt.Println(s1)
	conn.Write("PASS " + pw + "\r\n")
	s2 ,_ := conn.Read()
	fmt.Println(s2)
	
	conn.Write("LIST" + "\r\n")
	list ,_ := conn.Read()
	fmt.Println(list)
	
	conn.Write("RETR" + "\r\n")
	retr ,_ := conn.Read()
	fmt.Println(retr)

	conn.Close()
	
	return conn
}
/*
func Pop3List(conn *Connection) string{
	
	
	conn.Write("LIST" + "\r\n")
	list ,_ := conn.Read()
	fmt.Println(list)
	
	return "list"
}

func Pop3Retr(conn *Connection) string{

	conn.Write("RETR" + "\r\n")
	retr ,_ := conn.Read()
	fmt.Println(retr)
	
	return "retr"	
}
*/