package main

import (
	"Team5Project/connection"
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
		*pw = userInterface.GetPassword()
	}

	conn, err := connection.Pop3Auth("pop.nyx.net", "110", *un, *pw)
	defer conn.Close()
	if err == "err" {
		log.Fatal(err)
	}
	s := connection.Pop3List(conn)
	
	messages := connection.RetrieveAll(conn, connection.ExtractFromList(s))
	fmt.Println(messages)
	
	SaveConfig(*un, *pw)

}

func ParseConfig() (string, string) {
	config, err := ioutil.ReadFile(".config")
	if err != nil {
		log.Fatal(err)
	}
	conStr := string(config)
	lines := strings.Split(conStr, "\n")
	un := lines[0]
	pw := lines[1]
	return un, pw
}

/*
Config file layout:
username
password
*/
func SaveConfig(un, pw string) {
	file, err := os.OpenFile(".config", os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := []byte(fmt.Sprintf("%s\n%s", un, pw))
	file.Write(bytes)
}
