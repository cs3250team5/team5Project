package connection

import (
	"strings"
	"fmt"
	"io/ioutil"
	"os"
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

func check(e err){
	if e != nil{
		panic(e)
		}
	}
}

func (mail MailObject) save(){
	f, err := os.create("/temp/"+ mail.From +"_" + mail.Date)
	check(err)
	
	defer f.close()
	
	s = bufio.NewWriter(f)
	m = s.WriteString(mail.from + mail.Date + mail.To + mail.Subject)
	m = s.WriteString(mail.message)
	
	s.Flush()
	
}

func Read (file string)(MailObject){
	f, err := os.Open(file)
	check(err)
	defer f.close()
	
	var m MailObject
	m.From = f.ReadLine()
	m.Date = f.ReadLine()
	m.To = f.ReadLine()
	m.message = f.Read()
	return m
}