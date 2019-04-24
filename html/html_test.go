package html

import (
	"io/ioutil"
	"testing"
)

func TestPullText(t *testing.T) {
	sampleb, _ := ioutil.ReadFile("testmail.txt")
	sample := string(sampleb)
	PullText(sample)
}
