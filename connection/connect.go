package connect

import (
	"fmt"
	"net"
	"os"
)

type Connection struct{
	Con net.Conn
}

func(connect Connection) Open(port string, host string) {
	connection, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error in Client main: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected")
	connect.Con = connection
}

func(connect Connection) Close(){
	connect.Con.Close()
	fmt.Println("Connection Closed")
}


