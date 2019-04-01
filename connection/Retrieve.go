package connection

import(
	"fmt"
)

func RetrieveAll(conn *Connection, listMap map[int]int) map[int]MailObject {
	// Emails are put into maps based on data size
	finalMap := make(map[int]MailObject)
	for key, value := range listMap{
		s := Pop3Retr(conn, key, value)
		mail := MailFilter(s)
		finalMap[key] = mail
		fmt.Println("\nMessage", key, "\n", mail.To)
		
	}
	return finalMap
}