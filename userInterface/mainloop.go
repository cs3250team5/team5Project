package userInterface

import (
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
	}
}
