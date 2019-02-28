package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUsername() string {
	fmt.Print("Enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	un, _ := reader.ReadString('\n')
	return strings.TrimSpace(un)
}

func GetPassword() string {
	fmt.Print("Enter your password: ")
	reader := bufio.NewReader(os.Stdin)
	pw, _ := reader.ReadString('\n')
	return strings.TrimSpace(pw)
}

func GetUsernameAndPassword() (un, pw string) {
	un = GetUsername()
	pw = GetPassword()
	return
}
