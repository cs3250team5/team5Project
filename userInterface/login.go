package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUsername() string {
	
	// user tpyes username
	fmt.Print("Enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	un, _ := reader.ReadString('\n')
	return strings.TrimSpace(un)
}

func GetPassword() string {
	
	// User types Password
	fmt.Print("Enter your password: ")
	reader := bufio.NewReader(os.Stdin)
	pw, _ := reader.ReadString('\n')
	return strings.TrimSpace(pw)
}

func GetUsernameAndPassword() (un, pw string) {
	
	// Sets username and password to varible
	un = GetUsername()
	pw = GetPassword()
	return
}

func GetMessageNumber() string {
	
	// User types in Message number
	fmt.Print("Which message to read/download: ")
	reader := bufio.NewReader(os.Stdin)
	messNo, _ := reader.ReadString('\n')
	return strings.TrimSpace(messNo)
}
