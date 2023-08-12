package encrypt

import (
	"encoding/base64"
	"testing"

	"github.com/miteshbsjat/goshell"
)

func TestGenerateIV(t *testing.T) {
	randSeed := 1234
	line := "A quick brown fox jumps over the lazy dog."
	expectedIVB64 := "j6snQT7LLMdWEj8wA5b3gw=="
	iv, _ := generateIV([]byte(line), randSeed)
	if base64.StdEncoding.EncodeToString(iv) != expectedIVB64 {
		t.Errorf("Expected and generated IV does not match")
	}
}

func TestEncAlgos(t *testing.T) {
	initENCDECAlgosVar()
	if 0 == len(ENCRYPTION_ALGORITHMS) {
		t.Errorf("ENCRYPTION_ALGORITHMS[] is not initialized with encryptionFuncMap %d", len(ENCRYPTION_ALGORITHMS))
	}
	if len(encryptionFuncMap) != len(ENCRYPTION_ALGORITHMS) {
		t.Errorf("ENCRYPTION_ALGORITHMS[] is not set with encryptionFuncMap %v != %v", ENCRYPTION_ALGORITHMS, encryptionFuncMap)
	}
}

func TestEncryptionAES(t *testing.T) {

	plainFile := "/tmp/testenc.txt"
	_, _ = goshell.RunCommand("echo Hello World > " + plainFile)
	_, _ = goshell.RunCommand("echo '1' >> " + plainFile)
	_, _ = goshell.RunCommand("echo '1' >> " + plainFile)
	_, _ = goshell.RunCommand("echo Hello World >> " + plainFile)
	_, _ = goshell.RunCommand("echo '1' >> " + plainFile)
	_, _ = goshell.RunCommand("echo Hello World >> " + plainFile)

	encFile := plainFile + ".enc"
	passwd := []byte("passwordpassword")
	encFunc := encryptionFuncMap["aes"]
	err := EncryptFileLineByLine(plainFile, encFunc, passwd)
	if err != nil {
		t.Error(err)
	}
	decFunc := decryptionFuncMap["aes"]
	err = DecryptFileLineByLine(encFile, decFunc, passwd)
	if err != nil {
		t.Error(err)
	}

	decFile := encFile + ".dec"
	output, _ := goshell.RunCommand("head -1 " + decFile)
	if output != "Hello World" {
		t.Errorf("First line after decryption does not match")
	}
}
