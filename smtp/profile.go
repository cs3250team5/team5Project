package smtp

import(
	"fmt"
	"io/ioutil"
	"connection"
	"net/smtp"
	"regexp"
	
)

func EmailTo() string {
	
	// user writes message to 
	fmt.Print("To: ")
	reader := bufio.NewReader(os.Stdin)
	to, _ := reader.ReadString('\n')
	return strings.TrimSpace(to)
}

func EmailSubject() string {
	
	// user types subject of email
	fmt.Print("Subject: ")
	reader := bufio.NewReader(os.Stdin)
	sub, _ := reader.ReadString('\n')
	return strings.TrimSpace(sub)
}

func EmailMsg() string {
	
	// User types message
	fmt.Print("Write message")
	reader := bufio.NewReader(os.Stdin)
	writemsg, _ := reader.ReadString('\n')
	fmt.Print(".\n")
	return strings.TrimSpace(writemsg)
	clean = CleanMessage(writemsg)
	fmt.Print(strings.TrimSpace(clean))
}

}

func CompileMessage(EmailTo string, EmailSubject string, EmailMsg string){

	//creating the email
	sub := composeEmail.EmailSubject()
	writeMsg := composeEmail.EmailMsg()
	emailReciever := composeEmail.EmailTo()
	
	msg := []byte("To: " + emailReciever + "\r\n" + "Subject :" + sub + "\r\n" +  writeMsg)

	//user decides whether to save and send
	fmt.Print("Would you like save as a draft or send the message?// S/s D/d")
	fmt.Scanf("%s", &choice)
	if choice == 'S' || choice == 's'{
		msg = SendMail(msg)
	}
	if choice == 'D' || choice == 'd'{
		msg = draft(msg)
	}
}


func SendMail(conn *Connection, EmailTo string, EmailSubject string, EmailMsg string){
	
	//configuration
	hostURL := "smtp.gmail.com"//This will change
	hostPORT := "587"// This will change
	emailSender := "cs3250Team5Relay"
	password := "bdexlpeeudlnsuwy"

	//creating auth object
	emailAUTH := smtp.PlainAuth(
		"", emailSender, password, hostURL)

	//send the mail
	err := smtp.SendMail(
		hostURL+":"+hostPORT, emailAUTH, emailSender, []string{emailReciever}, msg)

	if err != nil {
		fmt.Print("Error :", err)
	}
	fmt.Println(" email sent from " + emailReciever)

	connect.Con.Close()
}

func draft(EmailTo string, EmailSubject string, EmailMsg string){
	
	var findLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	
	if findLetter == EmailTo || findLetter == EmailSubject || findLetter == EmailMsg{ 

	
		fileName := fmt.Sprintf("%d_%s_%s.txt", mail.Num, mail.Subject, cleanFrom(mail.From))
		fileName = strings.Replace(fileName, " ", "_", -1)
		dir, err := filepath.Abs("draft")
		check(err)

		f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Draft folder
		check(err)

		defer f.Close()
		d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, cleanDate(mail.Date), mail.Subject, mail.Message)
		f.Write([]byte(d))
		fmt.Print("Draft saved and draft folder made.")
	}
	
}
	
//mini function to clean the message
func CleanMessage(s string) string {
	s = strings.TrimSuffix(s, "\n.")
	return s
}
