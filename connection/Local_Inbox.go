package connection

import (	
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CheckInbox() bool {
	_, err := os.Stat("Inbox") //looking for path
	if os.IsNotExist(err) {    //path does not exist
		return false
	}
	return true
}

func CreateInbox() { //creates inbox dependant on boolean from CheckInbox()
	if !CheckInbox() {
		os.Mkdir("Inbox", os.ModePerm)
	}
}

func ReadInbox(inbox string) map[int]MailObject {
	files, err := ioutil.ReadDir(inbox)
	if err != nil {
		log.Fatal(err)
	}
	finalMap := make(map[int]MailObject /*may change*/)
	for _, file := range files {
		m := ReadMF(inbox + "/" + file.Name())
		finalMap[m.Num] = m
	}
	return finalMap
}

func WriteToInbox(m map[int]MailObject) {
	CreateInbox()
	for v := range m {
		Save(m[v])
	}
}

func Mail2Line(mail MailObject) string {
	str := fmt.Sprintf("|%4d|%20s|%30s|%10s",mail.Num,mail.From,mail.Subject,mail.Date)
	return str
}