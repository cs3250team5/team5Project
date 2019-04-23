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
	str1 := fmt.Sprintf("%.40s",mail.From)
	str2 := fmt.Sprintf("%.50s",mail.Subject)
	str3 := fmt.Sprintf("%.16s",mail.Date)
	str := fmt.Sprintf("|%-4d|%-40s|%-50s|%-16s|",mail.Num,str1,str2,str3)
	return str
}

func DisInbox(mail map[int]MailObject) {
	fmt.Println("```````````````````````````````````````````````````````````````````````````````````````````````````````````````````")
	for m := range mail{
		fmt.Printf(Mail2Line(mail[m]))
		fmt.Println()
		fmt.Println("```````````````````````````````````````````````````````````````````````````````````````````````````````````````````")
	}
}
