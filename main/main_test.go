package main

import (
	"Team5Project/connection"
	"Team5Project/smtp"
	"testing"
)

func TestDraft(t *testing.T) {
	var g connection.MailObject
	if smtp.SendMail(g, "email", "sub", "msg") != "no draft made" {

		t.Error("expected saved draft")

	}
}
