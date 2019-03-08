package connection

import (
	"fmt"
	"strings"
)

const msg = `X-Spam-Checker-Version: SpamAssassin 3.4.1 (2015-04-28) on hermes.nyx.net
X-Spam-Level:
X-Spam-Status: No, score=-0.1 required=5.0 tests=DKIM_SIGNED,DKIM_VALID,
        DKIM_VALID_AU,FREEMAIL_FROM,HTML_MESSAGE,RCVD_IN_DNSWL_NONE,RCVD_IN_MSPIKE_H2,
        SPF_PASS,TVD_SPACE_RATIO autolearn=disabled version=3.4.1
X-Spam-ASN: AS15169 209.85.128.0/17
X-Originating-IP: 209.85.167.49
Return-Path: <frederickbuker@gmail.com>
Received: from mail-lf1-f49.google.com (mail-lf1-f49.google.com [209.85.167.49])
        by hermes.nyx.net (8.15.2/8.15.2/fnord) ith ESMTPS id x1JMRKgj001073
        (version=TLSv1.2 cipher=ECDHE-RSA-AES128-GCM-SHA256 bits=128 verify=NOT)
        for <fbuker@nyx.net>; Tue, 19 Feb 2019 15:27:26 -0700
Received: by mail-lf1-f49.google.com with SMTP id l10so16022311lfh.9
        for <fbuker@nyx.net>; Tue, 19 Feb 2019 14:27:25 -0800 (PST)
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
        d=gmail.com; s=20161025;
        h=mime-version:from:date:message-id:subject:to;
        bh=xU659WFW9p3snoZrCygpn3AXtc9nLXWaWsPzodzl7R8=;
        b=I6VUlW4TIrVwAW0S0c05gmN6B7MJGHx1GEhDmC2y3hsb7v92ORW1H1MJ9s+DZiW84O
         hTy6hkKlimyK41NMiGgBUYVNxhK7Vrh53SWG36/M2LAIelGAezgImd0wH/KfauPyBLpX
         kB4YgvY4472lLJRlhdJcg4y5j4zGNSQ2BixbN86smiSxDBySIc/g0Kqc8WTK+pAvrn6l
         H2K1Ax9NOepJMCbE4dD39EmRuZ7l+stOdlCZy+uZnRp9VuQlnRYc0ZO8NMaY09KgwwWn
         J4UUE1r9E7m2MERWP1sgU96x2hv+Adf62hfq6kbm2avQwZmaFbcn4AZ4cN30j/2Zp1XU
         mkNA==
X-Google-DKIM-Si1; a=rsa-sha256; c=relaxed/relaxed;
        d=1e100.net; s=20161025;
        h=x-gm-message-state:mime-version:from:date:message-id:subject:to;
        bh=xU659WFW9p3snoZrCygpn3AXtc9nLXWaWsPzodzl7R8=;
        b=N52P+1tsmUPdRa8XV8cp/F2M7dGy5rYg+z2QOUwjgFkOJSTzOcvPQLr70KMnCeF8vX
         q+pJxTAYMom1mZkB/JG9LAvzH5lztezynI58YEv4R2L7UbeZytAV+n68g29CKjeROf4E
         DVUMGmCrdBTSucbBxcaKdla8isU1ENBU9/1dx7mtWl7Rz0fyhJkthk9QyF+1tzKwHcNu
         NjwmQzeRmCBRz864/9ttft5AX751RqYy9IeYo9Vhb/agLvOzzcZl6b7EEatTqp9DTpaG
         YP3Upxqx81ZIK7j3Q91yUlTXH3TcNsppHFG4kzv7nbTqsHp547chbhN/0FdVkhSBDjjY
         oo9A==
X-Gm-Message-State: AHQUAubZCA5R8whR0KVFvmKJ91ApOwp8n/4pEFcmXT3CXVI4KL+xO/3D
        AYJM2EmLrnQiS77C/DGj9zT++vXclg2MwfMqwW+ygQ==
X-Google-Smtp-Source: AHgI3Ibe1r/b6ErodtmqBToo5xrIcmsXDw8Wb9aTQiVWMhPzDT/PtC8nV8yQNckfCMCGAweYyB2yttGYRqpM9mcc4JQ=
X-Received: by 2002:ac2:5328:: with SMTP id f8mr15341749lfh.42.1550615239073;
 Tue, 19 Feb 2019 14:27:19 -0800 (PST)
MIME-Version: 1.0
From: Frederick Buker <frederickbuker@gmail.com>
Date: Tue, 19 Feb 2019 15:27:06 -0700
Message-ID: <CACGUUV+B-ACkt_4cngbXO3wvXvY2x1ar9U2A3=LVDCNovkwb9Q@mail.gmail.com>
Subject: test
To: fbuker@nyx.net
Content-Type: multipart/alternative; boundary="000000000000822316058246be95"
X-Greylist: Sender IP whitelisted, not delayed by milter-greylist-4.5.11 (hermes.nyx.net [172.31.7.85]); Tue, 19 Feb 2019 15:27:26 -0700 (MST)

--000000000000822316058246be95
Content-Type: text/plain; charset="UTF-8"

Message

--000000000000822316058246be95
Content-Type: text/html; charset="UTF-8"

<div dir="ltr">Message<div><br></div></div>

--000000000000822316058246be95--
.`

type MailObject struct {
	To      string
	From    string
	Date    string
	Subject string
	Message string
}

func main() {
	mail := ReadLines(msg)
	fmt.Printf("To: %s\nFrom: %s\nDate: %s\nSubject: %s\nMessage: %s\n", mail.To, mail.From, mail.Date, mail.Subject, mail.Message)
	a := []string{mail.Message}
	fmt.Print(a)
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
