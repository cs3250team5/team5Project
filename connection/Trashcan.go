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
	//Looks to see if there an Trash
	_, err := os.Stat("Trash") //looking for path
	if os.IsNotExist(err) {    //path does not exist
		return false
	}
	return true
}

func CreateTrash() { 
	//Creates Trash dependant on boolean from CheckTrash()
	if !CheckTrash() {
		os.Mkdir("Trash", os.ModePerm)
	}
}

func WriteToTrash(m map[int]MailObject) {
	//Fills Trash
	CreateTrash()
	for v := range m {
		Discard(m[v])
	}
}

func ReadTrash(trash string) map[int]MailObject {
	//Reads Trash
	files, err := ioutil.ReadDir(trash)
	if err != nil {
		log.Fatal(err)
	}
	finalMap := make(map[int]MailObject)
	for _, file := range files {
		m := ReadMF(trash + "/" + file.Name())
		finalMap[m.Num] = m
	}
	return finalMap
}

func Discard(mail MailObject) {
	//Discards email into Trash folder
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