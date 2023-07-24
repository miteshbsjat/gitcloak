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

func TestTextFileKVFactor(t *testing.T) {
	gckv, err := NewKVStore("ggcmap")
	if err != nil {
		t.Error(err)
	}

	gckv.Set("ab", "cd")
	ret, _ := gckv.Get("ab")
	if ret != "cd" {
		t.Errorf("%s != cd", ret)
	}

	gHash := "abcde"
	gcHash := "fghij"
	gckv.Set(gHash, gcHash)
	ret, _ = gckv.Get(gHash)
	if ret != gcHash {
		t.Errorf("%s != %s", ret, gcHash)
	}

	gckv1, err := NewKVStore("ggcmap")
	if err != nil {
		t.Error(err)
	}
	if gckv != gckv1 {
		t.Errorf("%v != %v\n", gckv, gckv1)
	}

	gckv.Delete("ab")
	gckv.Delete("abcde")
}
