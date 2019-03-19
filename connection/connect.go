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
	message, err := connect.ReadN(bytes)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	message = cleanInput(message)
	if lines := strings.Split(message, "\n"); lines[len(lines)-2] == "." {
		return message, nil
	}
	for {
		mess, err := connect.ReadN(bytes)
		if err != nil && err.Error() != "EOF" {
			return "", err
		}
		message = message + cleanInput(mess)
		if lines := strings.Split(message, "\n"); lines[len(lines)-2] == "." {
			return message, nil
		}
	}
	return message, nil
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
	buffer := make([]byte, n*2)
	numBytes, err := connect.Con.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	return cleanInput(string(buffer[:numBytes])), nil

}

func (connect *Connection) Close() {
	connect.Con.Close()
	fmt.Println("Connection Closed")
}

// Removes \r from incoming strings
func cleanInput(s string) string {
	return strings.Replace(s, "\r", "", -1)
}
