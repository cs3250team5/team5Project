package connection

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MailObject struct {
	To      string
	From    string
	Date    string
	Subject string
	Message string
}

func ReadLines(s string) MailObject {
	var mail MailObject
	var boundary string
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		first, rest := firstRest(line)
		if first == "To:" {
			mail.To = rest
		}
		if first == "From:" {
			mail.From = rest
		}
		if first == "Date:" {
			mail.Date = rest
		}
		if first == "Subject:" {
			mail.Subject = rest
		}
		if first == "Content-Type:" {
			boundary = pullQuote(line)
			mail.Message = fixBoundary(lines[i:], boundary)
			break
		}
	}
	return mail
}

func pullQuote(s string) string {
	message := ""
	quote := false
	for _, line := range s {
		if line == '"' {
			if quote == true {
				break
			} else {
				quote = true
			}
		} else if quote == true {
			message = message + string(line)
		}
	}
	return message
}

func fixBoundary(lines []string, boundary string) string {
	message := ""
	header, reader := false, false
	for i, line := range lines {
		first, _ := firstRest(line)
		if header == true {
			header = false
		} else if reader == true {
			if first == "--"+boundary {
				break
			}
			message = message + lines[i]
		} else if first == "--"+boundary {
			header = true
			reader = true
		}
	}
	return message
}

func firstRest(s string) (string, string) {
	if len(s) > 0 {
		var first string
		fields := strings.Fields(s)
		if len(fields) > 0 {
			first = fields[0]
		} else {
			return "", ""
		}
		var rest string
		if len(s) > len(first)+1 {
			rest = s[len(first)+1:]
		} else {
			return first, ""
		}
		return first, rest
	} else {
		return "", ""
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Save(mail MailObject) {

	f, err := os.Create("/temp/" + mail.From + "_" + mail.Date)
	check(err)
	defer f.close()

	s := bufio.NewWriter(f)
	m := s.WriteString("From: %s,\n", mail.From)
	m = s.WriteString("Date: %s,\n", mail.Date)
	m = s.WriteString("Subject: %s,\n", mail.Subject)
	m = s.WriteString("Message: %s,\n", mail.message)

	s.Flush()

	fmt.Println("Successfully saved email")

}

/*func Read(file string) MailObject {
	var m MailObject
	f, err := os.Open(file)
	check(err)
	defer file.close()



	m.From =
	m.Date =
	m.To =
	m.Subject =
	m.Message =
	return m
}
*/
