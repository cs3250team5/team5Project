package main

import(

	"testing"
	"Team5Project/smtp"
	"Team5Project/connection"
	
)

func TestDraft(t *testing.T){

	if draft(email, sub, msg) != "Draft saved and draft folder made."{
		
		t.Error("expected saved draft")
		
	}
}