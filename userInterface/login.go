package userInterface

import (
	"bufio"
	"fmt"
	"os"
)

func GetUsername() (un string) {
	fmt.Print("Enter thy username: ")
	reader := bufio.NewReader(os.Stdin)
	un, _ = reader.ReadString('\n')
	return
}

func GetPassword() (pw string) {
	fmt.Print("Enter thy password: ")
	reader := bufio.NewReader(os.Stdin)
	pw, _ = reader.ReadString('\n')
	return
}

func GetUsernameAndPassword() (un, pw string) {
	un = GetUsername()
	pw = GetPassword()
	return un, pw
}
