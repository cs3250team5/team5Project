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
	wmsg, _ := reader.ReadString('\n')
	fmt.Print(".\n")
	return strings.TrimSpace(emsg)
}

