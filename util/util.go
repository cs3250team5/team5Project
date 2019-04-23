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
