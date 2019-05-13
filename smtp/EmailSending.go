package smtp

import (
	"Team5Project/userInterface"
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"io/ioutil"
	"log"
)

type MailDraft struct {
	To, Subject, Message string
}

func ComposeSend(draft MailDraft, EmailTo string, EmailSubject string, EmailMsg string ){

	var defaultUN, defaultPW string

		if _, err := os.Stat(".config"); !os.IsNotExist(err) {
			defaultUN, defaultPW = ParseConfig()
		} else {
			defaultUN, defaultPW = "", ""
		}

	//Creating the email
	emailReciever := userInterface.EmailTo()
	sub := userInterface.EmailSubject()
	writeMsg := userInterface.EmailMsg()

	//Configuration
	hostURL := "smtp.gmail.com"
	hostPORT := "587"
	username := defaultUN
	password := defaultPW
	
	//Creating auth object
	emailAUTH := smtp.PlainAuth("", username, password, hostURL)
	msg := []byte("To: " + emailReciever + "\r\n" + "Subject :" + sub + "\r\n" + writeMsg)

	//User decides whether to save and send
	fmt.Print("Would you like save as a draft or send the message?// S/s D/d")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	fmt.Println(choice)

	//Send the mail
	if strings.HasPrefix(choice, "s") || strings.HasPrefix(choice, "S") {
		err := smtp.SendMail(hostURL+":"+hostPORT, emailAUTH, username, []string{emailReciever}, msg)
		if err != nil {
			fmt.Print("Error :", err)
		}
		fmt.Println("email sent to " + emailReciever)

	}
	//Draft mail
	if strings.HasPrefix(choice, "d") || strings.HasPrefix(choice, "D") {
		var findLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if findLetter(EmailTo) == true || findLetter(EmailSubject) == true || findLetter(EmailMsg) == true {
			// Puts drafts in draft folder
			fileName := fmt.Sprintf("%s.txt", sub)
			fileName = strings.Replace(fileName, " ", "_", -1)
			dir, err := filepath.Abs("draft")
			check(err)
			f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Draft folder
			check(err)
			defer f.Close()
			d := fmt.Sprintf("%s", msg)
			f.Write([]byte(d))
		}
	}
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

func CheckDraft() bool {
	//Looks to see if there an draft
	_, err := os.Stat("draft") //looking for path
	if os.IsNotExist(err) {    //path does not exist
		return false
	}
	return true
}

func CreateDraft() {
	//Creates draft dependant on boolean from CheckDraft()
	if !CheckDraft() {
		os.Mkdir("draft", os.ModePerm)
	}
}

func check(e error) {
	//Error check
	if e != nil {
		panic(e)
	}
}
