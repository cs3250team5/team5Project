package main

import (
	"fmt"
	"os"
)

func CheckInbox(path string) bool {
	_, err := os.Stat(path) //path does exist
	if err == nil {
		return true
	}
	if os.IsNotExist(err) { //path does not exist
		return false
	}
	return true
}

func CreateInbox(b bool) {
	if b == false {
		os.Mkdir("/Users/4theone/Documents/Inbox", os.ModePerm) //specific path in my local computer
	} else {
		fmt.Println("There is already an inbox")
	}

}

func main() {
	fmt.Println("Let's try to make a new folder")
	k := CheckInbox("/Users/4theone/Documents/Inbox") //specific path in my local computer
	CreateInbox(k)
	fmt.Println("well is there an inbox?", k)

}
