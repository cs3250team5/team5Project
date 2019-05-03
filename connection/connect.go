package connection

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

type Connection struct {
	Con tls.Conn
}

func MakeConnection(host string, port string) *Connection {
	//Make connection to gmail server
	var conn Connection
	conn.Open(host, port)
	return &conn
}
func (connect *Connection) Open(host string, port string) {
	// Checks the connection to server
	connection, err := tls.Dial("tcp", host+":"+port, nil)
	if err != nil {
		fmt.Println("Error in Client main: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected")
	connect.Con = *connection
}

func (connect *Connection) ReadLines(bytes int) (string, error) {
	// Checks errors
	message, err := connect.ReadN(bytes)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	// Split stirng and ignores last 2 lines
	message = cleanInput(message)
	if lines := strings.Split(message, "\n"); lines[len(lines)-2] == "." {
		return message, nil
	}
	// Checks errors and splits strings and ignores last 2 lines
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

func (connect *Connection) Read() (string, error) {
	// Makes Memory and checks errors
	buffer := make([]byte, 1024)
	numBytes, err := connect.Con.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	return cleanInput(string(buffer[:numBytes])), nil
}

func (connect *Connection) ReadN(n int) (string, error) {
	// Makes Memory and checks errors
	buffer := make([]byte, n*2)
	numBytes, err := connect.Con.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	return cleanInput(string(buffer[:numBytes])), nil
}

func (connect *Connection) Write(s string) {
	connect.Con.Write([]byte(s))
}

func (connect *Connection) Close() {
	connect.Con.Close()
	fmt.Println("Connection Closed")
}

// Removes \r from incoming strings
func cleanInput(s string) string {
	return strings.Replace(s, "\r", "", -1)
}
