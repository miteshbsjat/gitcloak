package gitcloak

import "testing"

func TestTextFileKV(t *testing.T) {
	tkv := GetTextFileKV()
	tkv.Set("a", "b")
	ret, _ := tkv.Get("a")
	if ret != "b" {
		t.Errorf("%s != b", ret)
	}

	gHash := "abcd"
	gcHash := "fghij"
	PutGitAndGitCloak(gHash, gcHash)
	ret, _ = GetGitCloakCommitHash(gHash)
	if ret != gcHash {
		t.Errorf("%s != %s", ret, gcHash)
	}

}
