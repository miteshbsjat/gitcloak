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
