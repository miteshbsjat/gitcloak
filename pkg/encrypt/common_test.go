package encrypt

import (
	"encoding/base64"
	"sync"
	"testing"

	"github.com/miteshbsjat/gitcloak/pkg/fs"
	"github.com/miteshbsjat/gitcloak/pkg/gitcloak"
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
	err := EncryptFileLineByLine(plainFile, encFile, encFunc, passwd, SEED_DEFAULT, false)
	if err != nil {
		t.Error(err)
	}
	decFile := encFile + ".dec"
	decFunc := decryptionFuncMap["aes"]
	err = DecryptFileLineByLine(encFile, decFile, decFunc, passwd)
	if err != nil {
		t.Error(err)
	}

	output, _ := goshell.RunCommand("head -1 " + decFile)
	if output != "Hello World" {
		t.Errorf("First line after decryption does not match")
	}
	// t.Errorf("failing ...")
	goshell.RunCommand("rm -f " + plainFile + " " + encFile + " " + decFile)
}

func TestEncryptionXXTEA(t *testing.T) {

	plainFile := "/tmp/testencxt.txt"
	_, _ = goshell.RunCommand("echo Hello World > " + plainFile)
	_, _ = goshell.RunCommand("echo '1' >> " + plainFile)
	_, _ = goshell.RunCommand("echo '1' >> " + plainFile)
	_, _ = goshell.RunCommand("echo Hello World >> " + plainFile)
	_, _ = goshell.RunCommand("echo '1' >> " + plainFile)
	_, _ = goshell.RunCommand("echo Hello World >> " + plainFile)

	encFile := plainFile + ".enc"
	passwd := []byte("passwordpassword")
	encFunc := encryptionFuncMap["xxtea"]
	err := EncryptFileLineByLine(plainFile, encFile, encFunc, passwd, SEED_DEFAULT, true)
	if err != nil {
		t.Error(err)
	}
	decFile := encFile + ".dec"
	decFunc := decryptionFuncMap["xxtea"]
	err = DecryptFileLineByLine(encFile, decFile, decFunc, passwd)
	if err != nil {
		t.Error(err)
	}

	output, _ := goshell.RunCommand("head -1 " + decFile)
	if output != "Hello World" {
		t.Errorf("First line after decryption does not match")
	}
	// t.Errorf("failing ...")
	goshell.RunCommand("rm -f " + plainFile + " " + encFile + " " + decFile)
}

func TestTraversalEncryption(t *testing.T) {
	rootDir := gitcloak.GetGitCloakBase() + "/.."
	// regexPattern := `.*_test.go$`
	regexPattern := `.*/testencrypt/.*mitesh.*.txt$`

	regex, err := fs.RegexFromPattern(regexPattern)
	if err != nil {
		t.Errorf("Invalid regex pattern: %v", err)
		return
	}

	fileChannel := make(chan string, 10)
	errorChannel := make(chan error)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go fs.FindMatchingFiles(rootDir, regex, fileChannel, errorChannel, &wg)
	encFunc := encryptionFuncMap["aes"]
	key := []byte("passwordpassword")
	go EncryptFiles(fileChannel, errorChannel, done, encFunc, key, 1234, false)

	wg.Wait()
	close(fileChannel)

	<-done

	// Non-blocking getting message from channel
	select {
	case err := <-errorChannel:
		t.Errorf("received error %v", err)
	default:
		t.Log("No error")
	}
	// t.Errorf("Failing ...")

	rootDir = gitcloak.GetGitCloakBase() + "/.."
	// regexPattern := `.*_test.go$`
	regexPattern = `.*/testencrypt/.*mitesh.*.txt$`

	regex, err = fs.RegexFromPattern(regexPattern)
	if err != nil {
		t.Errorf("Invalid regex pattern: %v", err)
		return
	}

	fileChannel = make(chan string, 10)
	errorChannel = make(chan error)
	done = make(chan bool)

	wg.Add(1)

	go fs.FindMatchingFiles(rootDir, regex, fileChannel, errorChannel, &wg)
	decFunc := decryptionFuncMap["aes"]
	// goshell.RunCommand("sleep 1")
	go DecryptFiles(fileChannel, errorChannel, done, decFunc, key)

	wg.Wait()
	close(fileChannel)

	<-done

	// Non-blocking getting message from channel
	select {
	case err := <-errorChannel:
		t.Errorf("received error %v", err)
	default:
		t.Log("No error")
	}
	// t.Errorf(regexPattern)
}
