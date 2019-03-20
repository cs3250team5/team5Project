package connection

import (
	"fmt"
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
	r := CheckInbox()
	if r == false {
		os.Mkdir("Inbox", os.ModePerm)
	} else {
		fmt.Println("There is already an inbox")
	}

}
