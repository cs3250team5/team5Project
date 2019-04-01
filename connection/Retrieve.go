package connection

import (
	"fmt"
)

func RetrieveAll(conn *Connection, listMap map[int]int) map[int]MailObject {

	finalMap := make(map[int]MailObject /*may change*/)
	for k, v := range listMap {

		s := Pop3Retr(conn, k, v)
		mail := InterpretLines(s)
		mail.Num = k
		finalMap[k] = mail
		fmt.Println("\nMessage", k, "\n", mail.To)

	}

	return finalMap
}
