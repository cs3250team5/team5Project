package connection

import ( 
	"fmt"
	"encoding/hex"
	)
	
func Pop3Auth(host string, port string, un string, pw string) {
	conn := MakeConnection(host, port)
	s, _ := conn.Read()
	fmt.Println(s)
	conn.Close()
}

func Interpret(mess string) string{
	var a string
	a, err =  hex.DecodeString(mess)
	if err != nil{
		panic(err)
	}
	return string(a)
	
	
}