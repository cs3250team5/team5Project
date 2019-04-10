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
	wmsg, _ := reader.ReadString('\n')
	fmt.Print(".\n")
	return strings.TrimSpace(emsg)
}



func SendMail(conn *Connection, EmailTo string, EmailSubject string, EmailMsg string){
	
	sub := composeEmail.EmailSubject()
	writeMsg := composeEmail.EmailMsg()

	//configuration
	hostURL := "smtp.gmail.com"//This will change
	hostPORT := "587"// This will change
	emailSender := "fbuker@nyx.net"
	password := "su$lpHUr.86097"
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

