package encrypt

import (
	"bufio"
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"os"
	"sort"

	. "github.com/miteshbsjat/gitcloak/pkg/utils"
)

// Package for common encryption functions

// Define a map to map the encryption function names to the actual functions
var encryptionFuncMap = map[string]func([]byte, []byte) (string, error){
	"xxtea": EncryptAES,
	"aes":   EncryptAES,
	// Add entries for other algorithms (e.g., "chacha" and "xxtea") if needed
}
var decryptionFuncMap = map[string]func([]byte, string) ([]byte, error){
	"aes": DecryptAES,
	// Add entries for other algorithms (e.g., "chacha" and "xxtea") if needed
}
var ENCRYPTION_ALGORITHMS []string

func init() {
	initENCDECAlgosVar()

}
func initENCDECAlgosVar() {
	ENCRYPTION_ALGORITHMS = make([]string, 0, len(encryptionFuncMap))
	for encFuncName := range encryptionFuncMap {
		ENCRYPTION_ALGORITHMS = append(ENCRYPTION_ALGORITHMS, encFuncName)
		fmt.Println(encFuncName)
	}
	sort.Strings(ENCRYPTION_ALGORITHMS)
}

func generateIV(line []byte, randomNumber int) ([]byte, error) {
	// Combine the line content, the seed, and the random number
	data := append(line, []byte(fmt.Sprintf("%d", randomNumber))...)

	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return nil, err
	}

	// Get the hash sum and truncate it to AES block size (16 bytes)
	hashSum := hash.Sum(nil)
	iv := make([]byte, aes.BlockSize)
	copy(iv, hashSum)

	return iv, nil
}

func EncryptFileLineByLine(filepath string, encryptionFunc func([]byte, []byte) (string, error), key []byte) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	encryptedFilePath := filepath + ".enc"
	encryptedFile, err := os.Create(encryptedFilePath)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Printf("%v", string(line))
		// lineCopy := make([]byte, len(line))
		// copy(lineCopy, line)

		encryptedLine, err := encryptionFunc(key, line)
		if err != nil {
			return err
		}

		_, err = encryptedFile.WriteString(encryptedLine + LineBreak())
		if err != nil {
			return err
		}
	}

	return nil
}
func DecryptFileLineByLine(filepath string, decryptionFunc func([]byte, string) ([]byte, error), key []byte) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	decryptedFilePath := filepath + ".dec"
	decryptedFile, err := os.Create(decryptedFilePath)
	if err != nil {
		return err
	}
	defer decryptedFile.Close()

	for scanner.Scan() {
		line := string(scanner.Bytes())

		decryptedLine, err := decryptionFunc(key, line)
		if err != nil {
			return err
		}
		// fmt.Println(decryptedLine)
		// decryptedLine = appen([]byte("\n"))

		_, err = decryptedFile.Write(decryptedLine)
		if err != nil {
			return err
		}
		_, err = decryptedFile.Write([]byte(LineBreak()))
		if err != nil {
			return err
		}
	}

	return nil
}
