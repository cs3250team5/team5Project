package smtp

import(
	"fmt"
	"io/ioutil"
	"connection"
	"net/smtp"
	
)

func EmailTo() string {
	
	// user writes message to 
	fmt.Print("To: ")
	reader := bufio.NewReader(os.Stdin)
	to, _ := reader.ReadString('\n')
	return strings.TrimSpace(to)
}

func EmailSubject() string {
	
	// user tpyes subject of email
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
}



func SendMail(conn *Connection, EmailTo string, EmailSubject string, EmailMsg string){
	
	sub := composeEmail.EmailSubject()
	writeMsg := composeEmail.EmailMsg()

	//configuration
	hostURL := "smtp.gmail.com"//This will change
	hostPORT := "587"// This will change
	emailSender := "cs3250Team5Relay"
	password := "bdexlpeeudlnsuwy"
	emailReciever := composeEmail.EmailTo()

	//creating auth object
	emailAUTH := smtp.PlainAuth(
		"", emailSender, password, hostURL)

	//creating the email
	msg := []byte("To: " + emailReciever + "\r\n" + "Subject :" + sub + "\r\n" +  writeMsg)

	//send the mail
	err := smtp.SendMail(
		hostURL+":"+hostPORT, emailAUTH, emailSender, []string{emailReciever}, msg)

	if err != nil {
		fmt.Print("Error :", err)
	}
	fmt.Println(" email sent from " + emailReciever)

	connect.Con.Close()
}

func draft(){
	
	fileName := fmt.Sprintf("%d_%s_%s.txt", mail.Num, mail.Subject, cleanFrom(mail.From))
	fileName = strings.Replace(fileName, " ", "_", -1)
	dir, err := filepath.Abs("Inbox")
	check(err)

	f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Inbox folder
	check(err)

	defer f.Close()
	d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, cleanDate(mail.Date), mail.Subject, mail.Message)
	f.Write([]byte(d))
}
	

