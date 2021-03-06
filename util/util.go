package util

import "strings"

func FirstRest(s string) (string, string) {
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
func CleanDate(s string) string {
	date := strings.Replace(ReadUntil(s, "-"), " ", "-", -1)
	date = strings.Replace(date, "-", " ", -1)
	return date
}

func Clean(s string) string {
	var reserved = [...]string{"/", "\\", "?", "%", "*", ":", "|", "\"", "<", ">", "."}
	str := s
	for _, c := range reserved {
		str = strings.Replace(str, c, "", -1)
	}
	return str
}

func ReadUntil(s, delim string) string {
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

func CleanFrom(s string) string {
	// Replace " " with "_"
	name := strings.Replace(ReadUntil(s, "<"), " ", "_", -1)
	name = strings.Replace(name, "\"", "", -1)
	return name
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
