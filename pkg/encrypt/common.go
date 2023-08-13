package encrypt

import (
	"bufio"
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"os"
	"sort"

	"github.com/miteshbsjat/gitcloak/pkg/fs"
	"github.com/miteshbsjat/gitcloak/pkg/gitcloak"
	. "github.com/miteshbsjat/gitcloak/pkg/utils"
)

// Package for common encryption functions
type encryptParams struct {
	key     []byte
	randInt int
}

// Define a map to map the encryption function names to the actual functions
var encryptionFuncMap = map[string]func(encryptParams, []byte) (string, error){
	"xxtea": EncryptXXTEA,
	"aes":   EncryptAES,
	// Add entries for other algorithms (e.g., "chacha" and "xxtea") if needed
}
var decryptionFuncMap = map[string]func([]byte, string) ([]byte, error){
	"xxtea": DecryptXXTEA,
	"aes":   DecryptAES,
	// Add entries for other algorithms (e.g., "chacha" and "xxtea") if needed
}
var ENCRYPTION_ALGORITHMS []string

func NewEncryptParams(key []byte, randInt int) *encryptParams {
	ep := encryptParams{
		key:     key,
		randInt: randInt,
	}
	return &ep
}

var SEED_DEFAULT = int64(0)

func NewEncryptParamsDefRandInt(key []byte) *encryptParams {
	ep := encryptParams{
		key:     key,
		randInt: 0,
	}
	return &ep
}

func init() {
	initENCDECAlgosVar()

}
func initENCDECAlgosVar() {
	ENCRYPTION_ALGORITHMS = make([]string, 0, len(encryptionFuncMap))
	for encFuncName := range encryptionFuncMap {
		ENCRYPTION_ALGORITHMS = append(ENCRYPTION_ALGORITHMS, encFuncName)
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

func EncryptFileLineByLine(filepath string, encryptedFilePath string, encryptionFunc func(encryptParams, []byte) (string, error), key []byte, seed int64, perLineRandom bool) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	encryptedFile, err := os.Create(encryptedFilePath)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	// Initialize the random number generator with the seed
	rng := getRandomNumberGenerator(seed ^ (fs.GetFilePathId(filepath, gitcloak.GITCLOAK_BASE)))
	encParams := NewEncryptParams(key, rng.Intn(10000))

	for scanner.Scan() {
		line := scanner.Bytes()
		// fmt.Printf("%v", string(line))
		if perLineRandom {
			encParams = NewEncryptParams(key, rng.Intn(10000))
		}
		encryptedLine, err := encryptionFunc(*encParams, line)
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

func EncryptFiles(fileChannel <-chan string, errorChannel chan<- error, done chan<- bool, encryptionFunc func(encryptParams, []byte) (string, error), key []byte, seed int64, perLineRandom bool) {
	for filename := range fileChannel {
		fmt.Println("Filename:", filename)
		err := EncryptFileLineByLine(filename, fs.EncryptedFilePattern(filename), encryptionFunc, key, seed, perLineRandom)
		if err != nil {
			Warn("Error: %v", err)
			errorChannel <- err
		}
	}

	done <- true
}

func DecryptFileLineByLine(filepath string, decryptedFilePath string, decryptionFunc func([]byte, string) ([]byte, error), key []byte) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
