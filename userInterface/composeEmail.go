package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EmailTo() string {

	// user writes message to
	fmt.Print("To: ")
	reader := bufio.NewReader(os.Stdin)
	to, _ := reader.ReadString('\n')
	return strings.TrimSpace(to)
}

func EmailSubject() string {

	// user tpyes subject of email
	fmt.Print("Subject: ")
	reader := bufio.NewReader(os.Stdin)
	sub, _ := reader.ReadString('\n')
	return strings.TrimSpace(sub)
}

func EmailMsg() string {
	// User types message
	var userInput string
	fmt.Print("Write message: ")
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line:= scanner.Text()
		if line =="."{ 
		break
		}
		userInput = userInput + line + "\n"

	}

	return userInput
}
