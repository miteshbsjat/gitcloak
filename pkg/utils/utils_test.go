package utils

import (
	"fmt"
	"testing"
)

func TestErrorExit(t *testing.T) {
	err := fmt.Errorf("Some Random Error %s", "GitCloak")
	if CheckIfError2(err, "Injecting error %s", "GitCloak") == false {
		t.Error(err)
	}
}

func TestLineBreak(t *testing.T) {
	lineBreak := LineBreak()
	if lineBreak != "\n" {
		t.Errorf("Line Break is not \\n")
	}
}
