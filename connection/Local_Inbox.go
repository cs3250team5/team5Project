package connection

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CheckInbox() bool {
	//Looks to see if there an inbox
	_, err := os.Stat("Inbox") //looking for path
	if os.IsNotExist(err) {    //path does not exist
		return false
	}
	return true
}

func CreateInbox() {
	//Creates inbox dependant on boolean from CheckInbox()
	if !CheckInbox() {
		os.Mkdir("Inbox", os.ModePerm)
	}
}

func ReadInbox(inbox string) map[int]MailObject {
	//Reads inbox
	files, err := ioutil.ReadDir(inbox)
	if err != nil {
		log.Fatal(err)
	}
	finalMap := make(map[int]MailObject)
	for _, file := range files {
		m := ReadMF(inbox + "/" + file.Name())
		finalMap[m.Num] = m
	}
	return finalMap
}

func WriteToInbox(m map[int]MailObject) {
	//Fill in inbox
	CreateInbox()
	for v := range m {
		Save(m[v])
	}
}

func ClearInbox(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
func Mail2Line(mail MailObject) string {
	//Adds summary of each mail for inbox
	str1 := fmt.Sprintf("%.40s", mail.From)
	str2 := fmt.Sprintf("%.50s", mail.Subject)
	str3 := fmt.Sprintf("%.16s", mail.Date)
	str := fmt.Sprintf("|%-4d|%-40s|%-50s|%-16s|", mail.Num, str1, str2, str3)
	return str
}

func DisInbox(mail map[int]MailObject) {
	//Makes Inbox look like and inbox
	fmt.Println("```````````````````````````````````````````````````````````````````````````````````````````````````````````````````")
	for m := range mail {
		fmt.Printf(Mail2Line(mail[m]))
		fmt.Println()
		fmt.Println("```````````````````````````````````````````````````````````````````````````````````````````````````````````````````")
	}
}
