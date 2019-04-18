package connection

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MailObject struct {
	To, From, Date, Subject, Message string
	Num                              int
}

func MailFilter(s string) MailObject {
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
	// Gatters all lines of message and adjust boundarys
	message := ""
	header, reader := false, false
	for i, line := range lines {
		first, _ := firstRest(line)
		if header == true {
			if line == "" {
				header = false
			}
		} else if reader == true {
			if first == "--"+boundary {
				break
			}
			message = message + "\n" + lines[i]
		} else if first == "--"+boundary {
			header = true
			reader = true
		}
	}
	return message
}

func firstRest(s string) (string, string) {
	// Looks at length of string
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
func cleanFrom(s string) string {
	// Replace " " with "_"
	name := strings.Replace(readUntil(s, "<"), " ", "_", -1)
	name = strings.Replace(name, "\"", "", -1)
	return name
}

func cleanDate(s string) string {
	date := strings.Replace(readUntil(s, "-"), " ", "-", -1)
	date = strings.Replace(date, "-", " ", -1)
	return date
}

func clean(s string) string {
	var reserved = [...]string{"/", "\\", "?", "%", "*", ":", "|", "\"", "<", ">", "."}
	str := s
	for _, c := range reserved {
		str = strings.Replace(str, c, "", -1)
	}
	return str
}

func readUntil(s, delim string) string {
	// Reads stirngs and gets rid of last line
	s1 := ""
	for _, c := range s {
		if string(c) != delim {
			s1 = s1 + string(c)
		} else {
			return s1[:len(s1)-1]
		}
	}
	return s1
}

func Save(mail MailObject) {
	// Saves emails
	fileName := fmt.Sprintf("%d-%s-%s.txt", mail.Num, clean(mail.Subject), cleanFrom(mail.From))
	fileName = strings.Replace(fileName, " ", "_", -1)
	dir, err := filepath.Abs("Inbox")
	check(err)

	f, err := os.Create(filepath.Join(dir, filepath.Base(fileName))) //creates file within local Inbox folder
	check(err)

	defer f.Close()
	d := fmt.Sprintf("Num: %d\nTo: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage:\n%s\n", mail.Num, mail.To, mail.From, cleanDate(mail.Date), mail.Subject, mail.Message)
	f.Write([]byte(d))
}

func ReadMF(file string) MailObject {
	// Make email interface better
	var m MailObject
	f, err := ioutil.ReadFile(file)
	check(err)
	str := string(f)
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "Num: ") {
			n := strings.TrimPrefix(line, "Num: ")
			in, err := strconv.Atoi(n)
			if err != nil {
				fmt.Print(err)
			}
			m.Num = in
		} else if strings.HasPrefix(line, "To: ") {
			m.To = strings.TrimPrefix(line, "To: ")
		} else if strings.HasPrefix(line, "From: ") {
			m.From = strings.TrimPrefix(line, "From: ")
		} else if strings.HasPrefix(line, "Date: ") {
			m.Date = strings.TrimPrefix(line, "Date: ")
		} else if strings.HasPrefix(line, "Subject: ") {
			m.Subject = strings.TrimPrefix(line, "Subject: ")
		} else if strings.HasPrefix(line, "Message:") {
			var s string
			for _, Mline := range lines[i+1:] {
				s = s + Mline + "\n"
			}
			m.Message = s
			break
		}
	}
	return m
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
