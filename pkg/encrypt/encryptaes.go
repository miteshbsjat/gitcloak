package encrypt

import (
	"crypto/aes"
	"crypto/cipher"

	// "crypto/rand"
	"encoding/base64"
	"errors"
	// "io"
)

func EncryptAES(encParams encryptParams, data []byte) (string, error) {
	key := encParams.key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not necessarily secure.
	// For now, let's assume we use a random IV for each line.
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	return "", err
	// }
	copy(ciphertext, data)
	ivg, _ := generateIV(ciphertext, encParams.randInt)
	for i := 0; i < aes.BlockSize; i++ {
		iv[i] = ivg[i]
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	// Encode the ciphertext in Base64 before returning
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return encodedCiphertext, nil
}

func DecryptAES(key []byte, encodedCiphertext string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
