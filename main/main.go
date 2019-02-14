package main

import (
	"Team5Project/userInterface"
	"fmt"
	"strings"
)

func main() {
	un, pw := userInterface.GetUsernameAndPassword()
	un = strings.Trim(un, " \r\n")
	pw = strings.Trim(pw, " \r\n")
	fmt.Printf("Username: %s Password: %s\n", un, pw)
}
