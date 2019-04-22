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
	fmt.Print("Write message")
	reader := bufio.NewReader(os.Stdin)
	writemsg, _ := reader.ReadString('\n')
	fmt.Print(".\n")
	clean := CleanMessage(writemsg)
	fmt.Print(strings.TrimSpace(clean))
	return strings.TrimSpace(writemsg)
}

//mini function to clean the message
func CleanMessage(s string) string {
	s = strings.TrimSuffix(s, "\n.")
	return s
}
