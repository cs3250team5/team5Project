package menu

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"Team5Project/connection"
	"Team5Project/smtp"
)

/*
states:
1 - compose message
2 - view inbox
*/

func RequestState() int {
	//Navigates through main menu
	for {
		fmt.Println("What would you like to do?\n1 - Write a Message\n2 - View Inbox\n")
		reader := bufio.NewReader(os.Stdin)
		res, _ := reader.ReadString('\n')
		resi, _ := strconv.Atoi(strings.TrimSpace(res))
		if resi == 1 || resi == 2 {
			return resi
		}
		fmt.Println("input not recognized\n Please enter 1 or 2\n")
	}
}

func InboxNavi() {
	//Nnavigates through inbox menu
	mails := connection.ReadInbox("Inbox")
	connection.DisInbox(mails)
	MailNavi(mails)
	/*
	fmt.Println("Open email? Y/N")
	reader := bufio.NewReader(os.Stdin)
	resi, _ := reader.ReadString('\n')
	for{
		if strings.HasPrefix(resi,"y")|| strings.HasPrefix(resi,"Y"){
			resi = MailNavi(mails)
		}
		if strings.HasPrefix(resi,"n")|| strings.HasPrefix(resi,"N"){
			break
		}else{
			fmt.Println("Not a valid input put in a valid input try again")	
			InboxNavi()// recusive for error handling
		}
	}*/
}

func MailNavi(mails map[int]connection.MailObject) string{
	//Lets user chose which email to read
	fmt.Println("Please enter Email Number:")
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	resi, _ := strconv.Atoi(strings.TrimSpace(res))
	connection.DisplayEmail(mails[resi])
	fmt.Println("Open another email? Y/N")
	res, _ = reader.ReadString('\n')
	if strings.HasPrefix(res,"y")|| strings.HasPrefix(res,"Y"){
			InboxNavi()
		}
	if strings.HasPrefix(res,"n")|| strings.HasPrefix(res,"N"){
		Menu()
		return res
	}
	mails = connection.ReadInbox("Inbox")
	connection.DisInbox(mails)
	return res
}

func Menu(){
	st := RequestState()
	
	if st == 2{
		InboxNavi()
	}
	if st == 1{
		var g smtp.MailDraft
		smtp.ComposeSend(g, "email", "sub", "msg")
	}
}