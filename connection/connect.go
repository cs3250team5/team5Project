package connection

import (
	"fmt"
	"net"
	"os"
)

type Connection struct {
	Con net.Conn
}

func (connect *Connection) Open(host string, port string) {
	connection, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error in Client main: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected")
	connect.Con = connection
}

func (connect *Connection) Write(s string) {
	fmt.Fprintf(connect.Con, s)
}

func (connect *Connection) Read() (string, error) {
	buffer := make([]byte, 1024)
	numBytes, err := connect.Con.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	return string(buffer[:numBytes]), nil
}

func (connect *Connection) Close() {
	connect.Con.Close()
	fmt.Println("*Connection Closed")
}
