package main

import (
	"Team5Project/connection"
	"Team5Project/userInterface"
	"fmt"
)

func main() {
	un, pw := userInterface.GetUsernameAndPassword()
	conn, err := connection.Pop3Auth("pop.nyx.net", "110", un, pw)
	defer conn.Close()
	if err != "err" {
		s := connection.Pop3List(conn)
		fmt.Println("Messages in Inbox:", s)
	}
}
