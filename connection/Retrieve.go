package connection

import (
	"fmt"
	"strconv"
)

func RetrieveAll(conn *Connection, listMap map[int]int) map[int]MailObject {

	finalMap := make(map[int]MailObject /*may change*/)
	for k, v := range listMap {

		s := Pop3Retr(conn, k, v)
		mail := InterpretLines(s)
		j := strconv.Itoa(k)
		mail.Key = j
		finalMap[k] = mail
		fmt.Println("\nMessage", k, "\n", mail.To)

	}

	return finalMap
}
