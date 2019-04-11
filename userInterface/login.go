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

func GetUsernameAndPassword() []string {
	fmt.Println("Is the login information the same for POP3 and SMTP?[y/n]")
	if yesNo() {
		un = GetUsername()
		pw = GetPassword()
		return []string{un, pw}
	}
	return nil
}

//Reads and interprets Yes/No Answers from the user. Does not ask questions
func yesNo() {
	reader := bufio.NewReader(os.Stdin)
	for {
		ans, _ := reader.ReadString("\n")
		if ans == 'y' || ans == 'n' {
			return ans
		} else {
			fmt.Println("Answer must be 'y' or 'n': ")
		}
	}
}
func GetMessageNumber() string {

	// User types in Message number
	fmt.Print("Which message to read/download: ")
	reader := bufio.NewReader(os.Stdin)
	messNo, _ := reader.ReadString('\n')
	return strings.TrimSpace(messNo)
}

func HelpText() string {
	helpText = "Team 5 Mail Client Usage:\n" +
		"Command line format:\n" +
		"\"commandline [FLAG [argument]]*\""
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
}
