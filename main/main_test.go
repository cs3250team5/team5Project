package main

import(

	"testing"	
)

func TestDraft(t *testing.T){

	if draft(email, sub, msg) != "Draft saved and draft folder made."{
		
		t.Error("expected saved draft")
		
	}
}