package main

import (
	"fmt"
	"strings"
)

func main() {

	msg := `Delivered-To: cs3250team5relay@gmail.com
    Received: by 2002:a4a:df50:0:0:0:0:0 with SMTP id j16csp1035854oou;
            Wed, 24 Apr 2019 13:33:13 -0700 (PDT)
    X-Received: by 2002:a63:ed4e:: with SMTP id m14mr33382819pgk.182.1556137992934;
            Wed, 24 Apr 2019 13:33:12 -0700 (PDT)
    ARC-Seal: i=1; a=rsa-sha256; t=1556137992; cv=none;
            d=google.com; s=arc-20160816;
            b=Q8zcB7LxUEzAVBUSIth2S1dov8UMJDZFQhWmT2WIVHLvgZSRiQ04SVbXb74bMMg5se
             3AHMu8jYjGR69J5WCFjepvBHD9W8V618juKVdKXPa6rbQtFMB5v1Hf/08ODMEYUMX6ZU
             Ae4MrcWB7DHZiFZ2uccgPg3goHLyYy3yaZ2X70dfaqkTcJp4cP1wQVywNCIfNhkrkXde
             hovHUFTBt4aCokLaSHZjnKBx5qW5J/zj+HzrHHH24OMdDVsTgIcQj8/7nZvtFIhohJgL
             hQ0qLXSQbfBhQgAWzEIoOY/inAAuanEPqGEhnPW5VZtKZsoO8FzqYZel222y1q+39rUp
             ZjiA==
    ARC-Message-Signature: i=1; a=rsa-sha256; c=relaxed/relaxed; d=google.com; s=arc-20160816;
            h=to:from:subject:message-id:feedback-id:date:mime-version
             :dkim-signature;
            bh=faYQBme2d8mkzM3eXjU224fT1H2eW5hEBME89ElUyVY=;
            b=eLj/AEZvnEuzkpsvoKZfcD4TIsoCm/J21J8Unuc5L/LTtmRHvAdnqGDyqHoaA3kLBk
             BiElxSPcQBk0SF6GIS8pJAvIEPw/MR77FfzCkJ/FuQZQTMrqNIUMxAPxJ7cIGkFTutPg
             42pgVg4oED2SwCnuMvlADgKzLHo3rxzsjh6ZiZ8BVhKydtpKFBq2Cq9TU4PXTNyQ5uLu
             7d9lMWODA/eqh2Gep2E72liyYVA1AElA5bHuQKdLF9aAnNkzAkoMY8GMXz0ZSAgHKPN6
             h/S3y+tyfpPQRi9jXM9Cqapwan5pO2d719W2XDPILU9P0rVABBRTwaSwDrDIyD/zSNMb
             sc6w==
    ARC-Authentication-Results: i=1; mx.google.com;
           dkim=pass header.i=@accounts.google.com header.s=20161025 header.b=sMqKd4Kg;
           spf=pass (google.com: domain of 3cmjaxagtaicyz-2p0w9lnnz5y43.rzzrwp.nzx@gaia.bounces.google.com designates 209.85.220.69 as permitted sender) smtp.mailfrom=3CMjAXAgTAIcyz-2p0w9lnnz5y43.rzzrwp.nzx@gaia.bounces.google.com;
           dmarc=pass (p=REJECT sp=REJECT dis=NONE) header.from=accounts.google.com
    Return-Path: <3CMjAXAgTAIcyz-2p0w9lnnz5y43.rzzrwp.nzx@gaia.bounces.google.com>
    Received: from mail-sor-f69.google.com (mail-sor-f69.google.com. [209.85.220.69])
            by mx.google.com with SMTPS id b8sor5518749pgr.43.2019.04.24.13.33.12
            for <cs3250team5relay@gmail.com>
            (Google Transport Security);
            Wed, 24 Apr 2019 13:33:12 -0700 (PDT)
    Received-SPF: pass (google.com: domain of 3cmjaxagtaicyz-2p0w9lnnz5y43.rzzrwp.nzx@gaia.bounces.google.com designates 209.85.220.69 as permitted sender) client-ip=209.85.220.69;
    Authentication-Results: mx.google.com;
           dkim=pass header.i=@accounts.google.com header.s=20161025 header.b=sMqKd4Kg;
           spf=pass (google.com: domain of 3cmjaxagtaicyz-2p0w9lnnz5y43.rzzrwp.nzx@gaia.bounces.google.com designates 209.85.220.69 as permitted sender) smtp.mailfrom=3CMjAXAgTAIcyz-2p0w9lnnz5y43.rzzrwp.nzx@gaia.bounces.google.com;
           dmarc=pass (p=REJECT sp=REJECT dis=NONE) header.from=accounts.google.com
    DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
            d=accounts.google.com; s=20161025;
            h=mime-version:date:feedback-id:message-id:subject:from:to;
            bh=faYQBme2d8mkzM3eXjU224fT1H2eW5hEBME89ElUyVY=;
            b=sMqKd4KgOX0Q+md4IV+tJDVgtMisoAaz5FI5+aBrtckpo3x2c8ajB/DeBz2UD1+pcM
             bHHs8EPF34jDC/aw6O3YnSAMLNsdgwgqsj6uD0LLcmKATZoU0DKoMiziUE5bIifYBj7Y
             uW50rfdbTOuVbeoo9rnpvRR5hpH8T7pyojIKZHjnNVMuKBhUSQuRKliPKoyOZesSQ1Ns
             Wqdu7jw2gDQ8JjopX1jPs9e8fdKAreCLXzul/FNgNi3Sgc37/YezF+Vl/lPCC28Ml4u6
             SCtjcKRjQycTPAnfxRShtsCLCYvB4VWaYoxp4+3Q+VY9rUkMcAH28qa8TG3U3n2GHzWW
             2gog==
    X-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
            d=1e100.net; s=20161025;
            h=x-gm-message-state:mime-version:date:feedback-id:message-id:subject
             :from:to;
            bh=faYQBme2d8mkzM3eXjU224fT1H2eW5hEBME89ElUyVY=;
            b=MqH6sEa1v+CyI9SK0crwX5qaeU2EAAeIfRGUng0w5lonROftqLHksvuyKn1SRSnhY7
             Xj3jO4t0c7rCQYTHtosV8k+J80nLVhTvFT6feWjjQptJY23zniCqnoDSCcEngYP1Xa3f
             g3Y+qWHOKT8uRivNM4/sHw6Jz/qH8ODOUXfT0PRb4kavueQRv5y8Hyr4CQVVNFqjL+wV
             wbxNxR8vbYv3bKR6Oz51+5JpBFVODa1QgFFeN/llQiXjZHGg++7/d8+2pgeGRlag882h
             +39mgKNepkW33AOqSAR+Fw4F8Vy2HH3daysVzy7I1uA/F0zse8CB16qQWDzTLTCzINFV
             sPTw==
    X-Gm-Message-State: APjAAAXvH0Ds2gpcN9srGke2AF6q5phr9jixjB1Zl30hfUPZm2cbho9p u1gu8YhaRIHAgwVTZAZlP4sL9j8Vg7GL
    X-Google-Smtp-Source: APXvYqzQPjAoA7qJ2vsNrdBYHvqfsqcuKBFddLkrUE711aiR3ONt6bvzbuMp9aOGz/uwdTpaLyt9/xVReY5sks9Ppie42A==
    MIME-Version: 1.0
    X-Received: by 2002:a63:cc0c:: with SMTP id x12mr31597275pgf.336.1556137992599; Wed, 24 Apr 2019 13:33:12 -0700 (PDT)
    Date: Wed, 24 Apr 2019 20:33:11 +0000 (UTC)
    X-Account-Notification-Type: 31-anexp#givab-fa--mdv2-fa--hsc-control_b--ivab-fa
    Feedback-ID: 31-anexp#givab-fa--mdv2-fa--hsc-control_b--ivab-fa:account-notifier
    X-Notifications: ddfcebadf3000000
    Message-ID: <BSfIAt4SCozqe4TM9hOYFQ.0@notifications.google.com>
    Subject: Security alert
    From: Google <no-reply@accounts.google.com>
    To: cs3250team5relay@gmail.com
    Content-Type: multipart/alternative; boundary="00000000000045389f05874c9c83"
    
    --00000000000045389f05874c9c83
    Content-Type: text/plain; charset="UTF-8"; format=flowed; delsp=yes
    Content-Transfer-Encoding: base64
    
    TmV3IGRldmljZSBzaWduZWQgaW4gdG8NCg0KDQpjczMyNTB0ZWFtNXJlbGF5QGdtYWlsLmNvbQ0K
    WW91ciBHb29nbGUgQWNjb3VudCB3YXMganVzdCBzaWduZWQgaW4gdG8gZnJvbSBhIG5ldyBNYWMg
    ZGV2aWNlLiBZb3UncmUNCmdldHRpbmcgdGhpcyBlbWFpbCB0byBtYWtlIHN1cmUgaXQgd2FzIHlv
    dS4NCkNoZWNrIGFjdGl2aXR5DQo8aHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tL0FjY291bnRD
    aG9vc2VyP0VtYWlsPWNzMzI1MHRlYW01cmVsYXlAZ21haWwuY29tJmNvbnRpbnVlPWh0dHBzOi8v
    bXlhY2NvdW50Lmdvb2dsZS5jb20vYWxlcnQvbnQvMTU1NjEzNzk5MTAwMD9yZm4lM0QzMSUyNnJm
    bmMlM0QxJTI2ZWlkJTNELTUxOTE1Mzk0ODc4NTU3NzcxNzMlMjZldCUzRDAlMjZhbmV4cCUzRGdp
    dmFiLWZhLS1tZHYyLWZhLS1oc2MtY29udHJvbF9iLS1pdmFiLWZhPg0KWW91IHJlY2VpdmVkIHRo
    aXMgZW1haWwgdG8gbGV0IHlvdSBrbm93IGFib3V0IGltcG9ydGFudCBjaGFuZ2VzIHRvIHlvdXIN
    Ckdvb2dsZSBBY2NvdW50IGFuZCBzZXJ2aWNlcy4NCsKpIDIwMTkgR29vZ2xlIExMQywgMTYwMCBB
    bXBoaXRoZWF0cmUgUGFya3dheSwgTW91bnRhaW4gVmlldywgQ0EgOTQwNDMsIFVTQQ0K
    --00000000000045389f05874c9c83
    Content-Type: text/html; charset="UTF-8"
    Content-Transfer-Encoding: quoted-printable
    
    <!DOCTYPE html><html lang=3Den><head><meta content=3D"email=3Dno" name=3D"f
    ormat-detection"><meta content=3D"date=3Dno" name=3D"format-detection"><sty
    le>.awl a {color: #FFFFFF; text-decoration: none;} .abml a {color: #000000;
     font-family: Roboto-Medium,Helvetica,Arial,sans-serif; font-weight: bold; 
    text-decoration: none;} .adgl a {color: rgba(0, 0, 0, 0.87); text-decoratio
    n: none;} .afal a {color: #b0b0b0; text-decoration: none;} @media screen an
    d (min-width: 600px) {.v2sp {padding: 6px 30px 0px;} .v2rsp {padding: 0px 1
    0px;}} @media screen and (min-width: 600px) {.mdv2rw {padding: 40px 40px;}}
     </style><link href=3D"//fonts.googleapis.com/css?family=3DGoogle+Sans" rel
    =3Dstylesheet type=3D"text/css"></head><body bgcolor=3D"#FFFFFF" style=3D"m
    argin: 0; padding: 0;"><table border=3D0 cellpadding=3D0 cellspacing=3D0 he
    ight=3D"100%" lang=3Den style=3D"min-width: 348px;" width=3D"100%"><Tbody><
    tr height=3D32 style=3D"height: 32px;"><td></td></tr><tr align=3Dcenter><td
    ><div itemscope itemtype=3D"//schema.org/EmailMessage"><div itemprop=3Dacti
    on itemscope itemtype=3D"//schema.org/ViewAction"><link href=3D"https://acc
    ounts.google.com/AccountChooser?Email=3Dcs3250team5relay@gmail.com&amp;cont
    inue=3Dhttps://myaccount.google.com/alert/nt/1556137991000?rfn%3D31%26rfnc%
    3D1%26eid%3D-5191539487855777173%26et%3D0%26anexp%3Dgivab-fa--mdv2-fa--hsc-
    control_b--ivab-fa" itemprop=3Durl><meta content=3D"Review Activity" itempr
    op=3Dname></div></div><table border=3D0 cellpadding=3D0 cellspacing=3D0 sty
    le=3D"padding-bottom: 20px;max-width: 516px;min-width: 220px;"><Tbody><tr><
    td style=3D"width: 8px;" width=3D8></td><td><div align=3Dcenter class=3Dmdv
    2rw style=3D"border-style: solid; border-width: thin; border-color:#dadce0;
     border-radius: 8px; padding: 40px 20px;"><img height=3D24 src=3D"https://w
    ww.gstatic.com/accountalerts/email/googlelogo_color_188x64dp.png" style=3D"
    width: 75px; height: 24px; margin-bottom: 16px;" width=3D75><div style=3D"f
    ont-family: &#39;Google Sans&#39;,Roboto,RobotoDraft,Helvetica,Arial,sans-s
    erif;border-bottom: thin solid #dadce0; color: rgba(0,0,0,0.87); line-heigh
    t: 32px; padding-bottom: 24px;text-align: center; word-break: break-word;">
    <div style=3D"font-size: 24px;">New device signed in&nbsp;to</div><table al
    ign=3Dcenter style=3D"margin-top:8px;"><Tbody><tr style=3D"line-height: nor
    mal;"><td align=3Dright style=3D"padding-right:8px;"><img height=3D20 src
    =3D"https://www.gstatic.com/accountalerts/email/anonymous_profile_photo.png
    " style=3D"width: 20px; height: 20px; vertical-align: sub; border-radius: 5
    0%;;" width=3D20></td><td><a style=3D"font-family: &#39;Google Sans&#39;,Ro
    boto,RobotoDraft,Helvetica,Arial,sans-serif;color: rgba(0,0,0,0.87); font-s
    ize: 14px; line-height: 20px;">cs3250team5relay@gmail.com</a></td></tr></Tb
    ody></table></div><div style=3D"font-family: Roboto-Regular,Helvetica,Arial
    ,sans-serif; font-size: 14px; color: rgba(0,0,0,0.87); line-height: 20px;pa
    dding-top: 20px; text-align: center;">Your Google Account was just signed i
    n to from a new Mac device. You're getting this email to make sure it was y
    ou.<div style=3D"padding-top: 32px;text-align: center;"><a href=3D"https://
    accounts.google.com/AccountChooser?Email=3Dcs3250team5relay@gmail.com&amp;c
    ontinue=3Dhttps://myaccount.google.com/alert/nt/1556137991000?rfn%3D31%26rf
    nc%3D1%26eid%3D-5191539487855777173%26et%3D0%26anexp%3Dgivab-fa--mdv2-fa--h
    sc-control_b--ivab-fa" link-id=3D"main-button-link" style=3D"font-family: &
    #39;Google Sans&#39;,Roboto,RobotoDraft,Helvetica,Arial,sans-serif; line-he
    ight: 16px; color: #ffffff; font-weight: 400; text-decoration: none;font-si
    ze: 14px;display:inline-block;padding: 10px 24px;background-color: #4184F3;
     border-radius: 5px; min-width: 90px;" target=3D"_blank">Check activity</a>
    </div></div></div><div style=3D"text-align: left;"><div style=3D"font-famil
    y: Roboto-Regular,Helvetica,Arial,sans-serif;color: rgba(0,0,0,0.54);font-s
    ize: 11px; line-height: 18px; padding-top: 12px; text-align: center;"><div>
    You received this email to let you know about important changes to your Goo
    gle Account and services.</div><div style=3D"direction: ltr;">&copy; 2019 G
    oogle LLC, <a class=3Dafal style=3D"font-family: Roboto-Regular,Helvetica,A
    rial,sans-serif;color: rgba(0,0,0,0.54);font-size: 11px; line-height: 18px;
     padding-top: 12px; text-align: center;">1600 Amphitheatre Parkway, Mountai
    n View, CA 94043, USA</a></div></div></div></td><td style=3D"width: 8px;" w
    idth=3D8></td></tr></Tbody></table></td></tr><tr height=3D32 style=3D"heigh
    t: 32px;"><td></td></tr></Tbody></table><img height=3D1 src=3D"https://noti
    fications.googleapis.com/email/t/AFG8qyX62IOrbAFwZym0kgyHxg9OWyId81y3YxXGIZ
    7Gf0i2RP2uX853IePilOiJR56-0Swnoj0y7Qiy4LtIx4p_PldLNn76rF5iFQ7lvdkb4Z2zgi-_X
    tdYLoc1GH2pRRh6HV4Ci1i4K0r4anM8wN6dIjh8PXqnVCrYC3SdUbazrbPsCJD83k7cKcEcg5kj
    wTaLFzMTGWnOI3OHJY7zxzu8XqG-3WCKypfevLjdaLGBj4zcUbWc1OClONlNrQP2tupevoBclg/
    a.gif" width=3D1></body></html>
    --00000000000045389f05874c9c83--`

	MailFilter(msg)
}

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
			mail.Message = fixBoundary(lines[i+1:], boundary)
			break
		}
	}
	fmt.Printf("%s\nFr%s\nDa%s\nSubje%s\n%s\n", mail.To, mail.From, mail.Date, mail.Subject, mail.Message)
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
		} else if first == "--"+boundary && strings.HasPrefix(lines[i+1], "Content-Type: text/html;") {
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
