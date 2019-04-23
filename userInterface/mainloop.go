package userInterface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"team5Project/connection"
)

/*
states:
1 - compose message
2 - view inbox
*/
func RequestState() int {
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

func InboxNavi(){
	for{
		mails := connection.ReadInbox("Inbox")
		connection.DisInbox(mails)
		fmt.Println("Open email? Y/N")
		reader := bufio.NewReader(os.Stdin)
		res, _ := reader.ReadString('\n')
		resi, _ := strconv.Atoi(strings.TrimSpace(res))
		
	}
}