package smtp

import (
	//"Team5Project/connection"
	"Team5Project/userInterface"
	//"Team5Project/util"
	"bufio"
	"fmt"
	"path/filepath"
	"net/smtp"
	"os"
	"regexp" 
	"strings"
)
type MailDraft struct {
	To, Subject, Message string
}


func ComposeSend(/*mail connection.MailObject,*/ draft MailDraft, EmailTo string, EmailSubject string, EmailMsg string) {

	//creating the email
	emailReciever := userInterface.EmailTo()
	sub := userInterface.EmailSubject()
	writeMsg := userInterface.EmailMsg()


	//configuration
	hostURL := "smtp.gmail.com" //This will change
	hostPORT := "587"           // This will change
	emailSender := "cs3250Team5Relay"
	password := "bdexlpeeudlnsuwy"
	
	//creating auth object
	emailAUTH := smtp.PlainAuth("", emailSender, password, hostURL)
	msg := []byte("To: " + emailReciever + "\r\n" + "Subject :" + sub + "\r\n" + writeMsg)

	//user decides whether to save and send
	fmt.Print("Would you like save as a draft or send the message?// S/s D/d")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	fmt.Println(choice)
	//send the mail
	if strings.HasPrefix(choice,"s") || strings.HasPrefix(choice,"S"){
		err := smtp.SendMail(hostURL+":"+hostPORT, emailAUTH, emailSender, []string{emailReciever}, msg)

		if err != nil {
			fmt.Print("Error :", err)
		}
		fmt.Println(" email sent to " + emailReciever)
	}
	if strings.HasPrefix(choice,"d") || strings.HasPrefix(choice,"D"){
		var findLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if findLetter(EmailTo) == true || findLetter(EmailSubject) == true || findLetter(EmailMsg) == true {

			fileName := fmt.Sprintf("%s.txt", "new draft")
			fileName = strings.Replace(fileName, " ", "_", -1)
			dir, err := filepath.Abs("draft")
			check(err)

			f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Draft folder
			check(err)

			defer f.Close()
			d := fmt.Sprintf("%s", msg)
			f.Write([]byte(d))
			//fmt.Println(msg, emailAUTH)
			/*
			if err != nil {
			fmt.Print("Error :", err)
			}*/
		
		}
	}
	/*connect.Con.Close()*/
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
