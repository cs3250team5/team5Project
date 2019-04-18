package util

import "testing"

func TestFirstRest(t *testing.T) {
	if first, rest := FirstRest("Help I'm Alive"); first != "Help" || rest != "I'm Alive" {
		t.Error("Error in FirstRest")
	}
}
