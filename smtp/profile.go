package smtp

import (
	"Team5Project/connection"
	"Team5Project/userInterface"
	//"Team5Project/util"
	//"bufio"
	"fmt"
	//"path/filepath"
	"net/smtp"
	//"os"
	//"regexp"
	//"strings"
)

func SendMail(mail connection.MailObject, EmailTo string, EmailSubject string, EmailMsg string) {

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
/*
	//user decides whether to save and send
	fmt.Print("Would you like save as a draft or send the message?// S/s D/d")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	fmt.Print(".\n")
	fmt.Println(choice)
	//send the mail
	if choice == "S" || choice == "s" {*/
		err := smtp.SendMail(
			hostURL+":"+hostPORT, emailAUTH, emailSender, []string{emailReciever}, msg)

		if err != nil {
			fmt.Print("Error :", err)
	}
	fmt.Println(" email sent to " + emailReciever)
	/*}
	if choice == "D" || choice == "d" {
		var findLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if findLetter(EmailTo) == true || findLetter(EmailSubject) == true || findLetter(EmailMsg) == true {

			fileName := fmt.Sprintf("%d_%s_%s.txt", mail.Num, mail.Subject, util.CleanFrom(mail.From))
			fileName = strings.Replace(fileName, " ", "_", -1)
			dir, err := filepath.Abs("draft")
			check(err)

			f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Draft folder
			check(err)

			defer f.Close()
			d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, util.CleanDate(mail.Date), mail.Subject, mail.Message)
			f.Write([]byte(d))
			g := ("Draft saved and draft folder made.")
			fmt.Println(msg, emailAUTH)
			return g/*

		}
	}*/

	/*connect.Con.Close()*/

}


func check(e error) {
	if e != nil {
		panic(e)
	}
}