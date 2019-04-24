package main

import (
	"Team5Project/connection"
	"Team5Project/smtp"
	"testing"
)

func TestDraft(t *testing.T) {
	var g connection.MailObject
	if smtp.SendMail(g, "email", "sub", "msg") != "Draft saved and draft folder made." {
		

		t.Error("expected saved draft")

	}
}
