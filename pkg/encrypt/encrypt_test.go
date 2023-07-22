package encrypt

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	// fmt.Println("!! Shree Ganeshay Namah !!")

	data := []byte("Mitesh Singh Jat")
	key := []byte("somepasssomepass")
	enc, err := Encrypt(data, 1, "xxtea", key)
	if err != nil {
		t.Error(err)
	}

	dec, err := Decrypt(enc, "xxtea", key)
	if err != nil {
		t.Error(err)
	}
	if string(data) != string(dec) {
		t.Errorf("%s != %s", data, dec)
	}

}
