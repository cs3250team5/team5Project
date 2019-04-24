package html

//HTML parser to interprete messages
// based on examples for net/html package
// found at: https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
import (
	"fmt"
	"golang.org/x/net/html"
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

func PullText(s string) {
	//	read := strings.NewReader(s)
	//	z := html.NewTokenizer(read)
	//	for {
	//		tt := z.Next()

	//		switch {
	//		case tt == html.ErrorToken:
	//			return
	//		case tt == html.TextToken:
	//			fmt.Println("Found Text Token")
	//			t := z.Token()
	//			mess := t.String()
	//			fmt.Println(mess)
	//			fmt.Println(string(z.Text()))
	//			if mess != "" || strings.TrimSpace(mess) == "" {
	//				message = message + "\n" + mess
	//			}
	//		}
	//
	//	}
	/*
		links := []string{}
		doc, err := html.Parse(strings.NewReader(s))
		if err != nil {
			log.Fatal(err)
		}
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						links = append(links, a.Val)
						break
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
		return links
	*/
	domDocTest := html.NewTokenizer(strings.NewReader(s))
	previousStartTokenTest := domDocTest.Token()
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
			if len(TxtContent) > 0 && !strings.HasPrefix(TxtContent, "=") {
				fmt.Printf("%s ", TxtContent)
			}
		}
	}
}
