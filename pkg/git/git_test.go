package git

import (
	"testing"
)

func TestGitBaseDir(t *testing.T) {

	empty := ""
	result, err := GetGitBaseDir()
	if err != nil {
		t.Error(err)
	}

	if result == empty {
		t.Errorf("%s == %s", result, empty)
	}
	result1, err := GetGitBaseDir()
	if err != nil {
		t.Error(err)
	}
	if result == empty {
		t.Errorf("%s == %s", result, empty)
	}
	if result != result1 {
		t.Errorf("%s == %s", result, result1)
	}
}
