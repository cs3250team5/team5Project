package main

import (
	"Team5Project/smtp"
	"testing"
)

func TestDraft(t *testing.T) {

	if smtp.Draft(email, sub, msg) != "Draft saved and draft folder made." {

		t.Error("expected saved draft")

	}
}