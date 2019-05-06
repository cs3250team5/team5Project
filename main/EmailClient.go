package main

import (
	"Team5Project/connection"
	"Team5Project/menu"
	"Team5Project/userInterface"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var defaultUN, defaultPW string

	if _, err := os.Stat(".config"); !os.IsNotExist(err) {
		defaultUN, defaultPW = ParseConfig()
	} else {
		defaultUN, defaultPW = "", ""
	}

	un := flag.String("un", defaultUN, "Username")
	pw := flag.String("pw", defaultPW, "Password")
	flag.Parse()

	if *un == "" {
		*un = userInterface.GetUsername()
	}
	if *pw == "" {
		*pw = userInterface.GetAppPassword()
	}

	conn, err := connection.Pop3Auth("pop.gmail.com", "995", *un, *pw)
	defer conn.Close()

	if err == "err" {
		log.Fatal(err)
	}
	s := connection.Pop3List(conn)
	messages := connection.RetrieveAll(conn, connection.ExtractFromList(s))
	fmt.Println("Messages Downloaded: ", len(messages))
	connection.WriteToInbox(messages)
	SaveConfig(*un, *pw)
	menu.Menu()
	/*
		st := userInterface.RequestState()

		if st == 2{
			userInterface.InboxNavi()
		}
		if st == 1{
			var g smtp.MailDraft
			smtp.ComposeSend(g, "email", "sub", "msg")
		}*/
}

func ParseConfig() (string, string) {
	//Configs your username and password
	config, err := ioutil.ReadFile(".config")

	if err != nil {
		log.Fatal(err)
	}

	conStr := string(config)
	lines := strings.Split(conStr, "\n")
	un := strings.TrimSpace(lines[0])
	pw := strings.TrimSpace(lines[1])
	return un, pw
}

/*
Config file layout:
username
password
*/

func SaveConfig(un, pw string) {
	//Saves your configs
	file, err := os.OpenFile(".config", os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes := []byte(fmt.Sprintf("%s\n%s", un, pw))
	file.Write(bytes)
}
