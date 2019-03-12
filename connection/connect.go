package connection

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Connection struct {
	Con net.Conn
}

func MakeConnection(host string, port string) *Connection {
	var conn Connection
	conn.Open(host, port)
	return &conn
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

func (connect *Connection) ReadLines(bytes int) (string, error) {
	buffer := make([]byte, bytes)
	var s string
	numBytes, err := connect.Con.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	s = cleanInput(string(buffer[:numBytes]))
	lines := strings.Split(s, "\n")
	fmt.Println(lines)
	if lines[len(lines)-2] == "." {
		return s, nil
	}
	fmt.Println("Didn't return")
	return s, nil
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
	return cleanInput(string(buffer[:numBytes])), nil
}

func (connect *Connection) ReadN(n int) (string, error) {
	buffer := make([]byte, n)
	numBytes, err := connect.Con.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	return string(buffer[:numBytes]), nil

}

func (connect *Connection) Close() {
	connect.Con.Close()
	fmt.Println("Connection Closed")
}

// Removes \r from incoming strings
func cleanInput(s string) string {
	return strings.Replace(s, "\r", "", -1)
}
