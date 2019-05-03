package connection

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	)
	
	
func CheckTrash() bool {
	_, err := os.Stat("Trash") //looking for path
	if os.IsNotExist(err) {    //path does not exist
		return false
	}
	return true
}

func CreateTrash() { //creates inbox dependant on boolean from CheckInbox()
	if !CheckTrash() {
		os.Mkdir("Trash", os.ModePerm)
	}
}

func WriteToTrash(m map[int]MailObject) {
	CreateTrash()
	for v := range m {
		Discard(m[v])
	}
	
}

func ReadTrash(trash string) map[int]MailObject {
	files, err := ioutil.ReadDir(trash)
	if err != nil {
		log.Fatal(err)
	}
	finalMap := make(map[int]MailObject /*may change*/)
	for _, file := range files {
		m := ReadMF(trash + "/" + file.Name())
		finalMap[m.Num] = m
	}
	return finalMap
}

func Discard(mail MailObject) {
	fileName := fmt.Sprintf("%d-%s-%s.txt", mail.Num, clean(mail.Subject), cleanFrom(mail.From))
	fileName = strings.Replace(fileName, " ", "_", -1)
	dir, err := filepath.Abs("Trash")
	check(err)

	f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Inbox folder
	check(err)

	defer f.Close()
	d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, cleanDate(mail.Date), mail.Subject, mail.Message)
	f.Write([]byte(d))
}