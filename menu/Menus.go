package menu

import (
	"Team5Project/connection"
	"Team5Project/smtp"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
states:
1 - compose message
2 - view inbox
3 - quit
*/

func RequestState() int {
	//Navigates through main menu
	for {
		fmt.Println("What would you like to do?\n1 - Write a Message\n2 - View Inbox\n3 - Quit")
		reader := bufio.NewReader(os.Stdin)
		res, _ := reader.ReadString('\n')
		resi, _ := strconv.Atoi(strings.TrimSpace(res))
		if resi == 1 || resi == 2 || resi == 3 {
			return resi
		}
		fmt.Println("Input not recognized\n Please enter 1, 2, or 3\n")
	}
}

func InboxNavi() {
	//Navigates through inbox menu
	mails := connection.ReadInbox("Inbox")
	connection.DisInbox(mails)
	MailNavi(mails)
}

func MailNavi(mails map[int]connection.MailObject) string {
	//Lets user chose which email to read
	fmt.Println("Please enter Email Number:")
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	resi, _ := strconv.Atoi(strings.TrimSpace(res))
	connection.DisplayEmail(mails[resi])
	fmt.Println("Open another email? Y/N")
	res, _ = reader.ReadString('\n')
	if strings.HasPrefix(res, "y") || strings.HasPrefix(res, "Y") {
		InboxNavi()
	}
	if strings.HasPrefix(res, "n") || strings.HasPrefix(res, "N") {
		return res
	}
	mails = connection.ReadInbox("Inbox")
	connection.DisInbox(mails)
	return res
}

func Menu() {
	for {
		st := RequestState()

		if st == 2 {
			InboxNavi()
		}
		if st == 1 {
			var g smtp.MailDraft
			smtp.ComposeSend(g, "email", "sub", "msg")
		}
		if st == 3 {
			break
		}
	}
}
