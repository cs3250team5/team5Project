package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUsername() string {
	//User types username
	fmt.Print("Enter your GMAIL username: ")
	reader := bufio.NewReader(os.Stdin)
	un, _ := reader.ReadString('\n')
	return strings.TrimSpace(un)
}

func GetAppPassword() string {
	//User types Password
	fmt.Print("Enter your GMAIL App password: ")
	reader := bufio.NewReader(os.Stdin)
	pw, _ := reader.ReadString('\n')
	return strings.TrimSpace(pw)
}

func GetUsernameAndPassword() []string {
	un := GetUsername()
	pw := GetPassword()
	return []string{un, pw}
}

func yesNo() bool {
	//Reads and interprets Yes/No Answers from the user. Does not ask questions
	reader := bufio.NewReader(os.Stdin)
	for {
		ans, _ := reader.ReadString('\n')
		fmt.Println(ans)
		if ans == "y" {
			return true
		} else if ans == "n" {
			return false
		} else {
			fmt.Println("Answer must be 'y' or 'n': ")
		}
	}
}

func GetMessageNumber() string {
	//User types in Message number
	fmt.Print("Which message to read/download: ")
	reader := bufio.NewReader(os.Stdin)
	messNo, _ := reader.ReadString('\n')
	return strings.TrimSpace(messNo)
}

func HelpText() string {
	//Helpful text for user
	helpText := "Team 5 Mail Client Usage:\n" +
		"Command line format:\n" +
		"\"commandline [FLAG [argument]]*\"" +
		"FLAGS:       Function:\n" +
		"-un          Changes the Username across all scopes (pop/smtp)\n" +
		"-pw          Changes the Password across all scope\n\n" +
		"-p3un        Changes the Pop3 Username\n" +
		"-p3pw        Changes the Pop3 Password\n\n" +
		"-smtpun/-su  Changes the SMTP Username\n" +
		"-smtppw/-sp  Changes the SMTP Password\n\n" +
		"-host/-h     Changes the Host across all scopes\n" +
		"-phost/-ph   Changes the Pop3 Host\n" +
		"-shost/-sh   Changes the SMTP Host\n\n"
	return helpText
}
