package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EmailTo() string {
	//User writes message to
	fmt.Print("To: ")
	reader := bufio.NewReader(os.Stdin)
	to, _ := reader.ReadString('\n')
	return strings.TrimSpace(to)
}

func EmailSubject() string {
	//User tpyes subject of email
	fmt.Print("Subject: ")
	reader := bufio.NewReader(os.Stdin)
	sub, _ := reader.ReadString('\n')
	return strings.TrimSpace(sub)
}

func EmailMsg() string {
	//User types message
	var userInput string
	fmt.Print("Write message: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "." {
			break
		}
		userInput = userInput + line + "\n"
	}
	return userInput
}
