package main

import (
	"Team5Project/connection"
	"fmt"
)

func main() {
	connect := connection.MakeConnection("localhost", "8080")

	connect.Write("Message")

	s, _ := connect.Read()
	fmt.Printf(s)
}
