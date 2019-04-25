package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"Team5Project/connection"
)

/*
states:
1 - compose message
2 - view inbox
*/
func RequestState() int {
	for {
		fmt.Println("What would you like to do?\n1 - Write a Message\n2 - View Inbox\n3-Close Program")
		reader := bufio.NewReader(os.Stdin)
		res, _ := reader.ReadString('\n')
		resi, _ := strconv.Atoi(strings.TrimSpace(res))
		if resi == 1 || resi == 2 || resi == 3{
			return resi
		}
			fmt.Println("input not recognized\n Please enter 1, 2, or 3\n")
	}
}

func InboxNavi(){
	mails := connection.ReadInbox("Inbox")
	connection.DisInbox(mails)
	fmt.Println("Open email? Y/N")
	reader := bufio.NewReader(os.Stdin)
	resi, _ := reader.ReadString('\n')
	for{
		if strings.HasPrefix(resi,"y")|| strings.HasPrefix(resi,"Y"){
			resi = MailNavi(mails)
		}
		if strings.HasPrefix(resi,"n")|| strings.HasPrefix(resi,"N"){
			break
		}			
	}
}

func MailNavi(mails map[int]connection.MailObject) string{
	fmt.Println("Please enter Email Number:")
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	resi, _ := strconv.Atoi(strings.TrimSpace(res))
	connection.DisEmail(mails[resi])
	fmt.Println("Open another email? Y/N")
	res, _ = reader.ReadString('\n')
	if strings.HasPrefix(res,"n")|| strings.HasPrefix(res,"N"){
			return res
	}
	mails = connection.ReadInbox("Inbox")
	connection.DisInbox(mails)
	return res
}

