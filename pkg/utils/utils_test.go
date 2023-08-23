package utils

import (
	"testing"
)

// func TestErrorExit(t *testing.T) {
// 	err := fmt.Errorf("Some Random Error %s", "GitCloak")
// 	CheckIfError2(err, "Injecting error %s", "GitCloak")
// }

func TestLineBreak(t *testing.T) {
	lineBreak := LineBreak()
	if lineBreak != "\n" {
		t.Errorf("Line Break is not \\n")
	}
}
