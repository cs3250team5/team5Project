package smtp

import (
	"Team5Project/connection"
	"Team5Project/userInterface"
	"Team5Project/util"
	"bufio"
	"fmt"
	//"io/ioutil"
	"net/smtp"
	"os"
	"regexp"
	"strings"
)
/*
func CompileMessage(EmailTo string, EmailSubject string, EmailMsg string) {

	//creating the email
	sub := userInterface.EmailSubject()
	writeMsg := userInterface.EmailMsg()
	emailReciever := userInterface.EmailTo()

	msg := []byte("To: " + emailReciever + "\r\n" + "Subject :" + sub + "\r\n" + writeMsg)
	
	fmt.Println(msg)

	//user decides whether to save and send
	fmt.Print("Would you like save as a draft or send the message?// S/s D/d")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	if choice == "S" || choice == "s" {
		msg := SendMail(emailReciever, sub, writeMsg)
	}
	if choice == "D" || choice == "d" {
		msg := Draft(msg)
	}
	msg := Draft(emailReciever, sub, writeMsg)
	return msg
}
*/

func SendMail(conn *Connection, EmailTo string, EmailSubject string, EmailMsg string, mail.MailObject) {

	//creating the email
	sub := userInterface.EmailSubject()
	writeMsg := userInterface.EmailMsg()
	emailReciever := userInterface.EmailTo()

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
	
	//send the mail
	if choice == "S" || choice == "s" {
		err := smtp.SendMail(hostURL+":"+hostPORT, emailAUTH, emailSender, []string{emailReciever}, msg)
		if err != nil {
			fmt.Print("Error :", err)
		}
		fmt.Println(" email sent from " + emailReciever)
	}
	if choice == "D" || choice == "d" {
			var findLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if findLetter(EmailTo) == true || findLetter(EmailSubject) == true || findLetter(EmailMsg) == true {

			fileName := fmt.Sprintf("%d_%s_%s.txt", conn.mail.Num, conn.mail.Subject, util.CleanFrom(mail.From))
			fileName = strings.Replace(fileName, " ", "_", -1)
			dir, err := filepath.Abs("draft")
			check(err)

			f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Draft folder
			check(err)

			defer f.Close()
			d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, cleanDate(mail.Date), mail.Subject, mail.Message)
			f.Write([]byte(d))
			g := ("Draft saved and draft folder made.")
		}
	}
	/*connect.Con.Close()*/
}

/*
func Draft(EmailTo string, EmailSubject string, EmailMsg string) string {

	var findLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if findLetter == EmailTo || findLetter == EmailSubject || findLetter == EmailMsg {

		fileName := fmt.Sprintf("%d_%s_%s.txt", mail.Num, mail.Subject, cleanFrom(mail.From))
		fileName = strings.Replace(fileName, " ", "_", -1)
		dir, err := filepath.Abs("draft")
		check(err)

		f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Draft folder
		check(err)

		defer f.Close()
		d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, cleanDate(mail.Date), mail.Subject, mail.Message)
		f.Write([]byte(d))
		g := ("Draft saved and draft folder made.")
		return g
	}

}

*/