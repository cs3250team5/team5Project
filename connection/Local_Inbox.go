package connection

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func CheckInbox() bool {
	_, err := os.Stat("Inbox") //looking for path
	if os.IsNotExist(err) {    //path does not exist
		return false
	}
	return true
}

func CreateInbox() { //creates inbox dependant on boolean from CheckInbox()
	r := CheckInbox()
	if r == false {
		os.Mkdir("Inbox", os.ModePerm)
	} else {
		fmt.Println("There is already an inbox")
	}

}

func ReadInbox(inbox string) map[int]MailObject{
	files, err := ioutil.ReadDir(inbox)
	if err != nil{
		log.Fatal(err)
	}
	finalMap := make(map[int]MailObject /*may change*/)
	for _, file := range files{
		m := ReadMF(file.Name())
		finalMap[m.Num] = m
	}
	return finalMap
}