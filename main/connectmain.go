package main

import (
	"Team5Project/connection"
	"fmt"
)

func main() {
	var connection connection.Connection
	connection.Open("localhost", "8080")

	connection.Write("Message")

	s, _ := connection.Read()
	fmt.Printf(s)
}
