package html

//HTML parser to interprete messages
// based on examples for net/html package
// found at: https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

func ReadHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func PullText(s string) string {
	links := ""
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = links + a.Val + "\n"
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	domDocTest := html.NewTokenizer(strings.NewReader(s))
	previousStartTokenTest := domDocTest.Token()
	ret := ""
loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))

			if previousStartTokenTest.Data == "title" {
				ret = ret + fmt.Sprintf("%s\n", TextDecoder(TxtContent))
			} else if len(TxtContent) > 0 && !strings.HasPrefix(TxtContent, "@media") {
				ret = ret + fmt.Sprintf("%s ", TextDecoder(TxtContent))
			}
		}
	}
	return ret + "\n--------------------------------------\nLinks In Mail:\n" + links
}

func TextDecoder(s string) string {
	s = strings.Replace(s, "=20", " ", -1)
	s = strings.Replace(s, "=09", "", -1)
	s = strings.Replace(s, "=E2=80=99", "'", -1)
	s = strings.Replace(s, "=C2=A9", "©", -1)
	return s
}
