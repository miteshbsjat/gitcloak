package encrypt

import (
	"crypto/aes"

	// "crypto/rand"

	// "io"
	"github.com/yang3yen/xxtea-go/xxtea"
)

func EncryptXXTEA(encParams encryptParams, data []byte) (string, error) {
	key := encParams.key

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	return "", err
	// }
	// copy(ciphertext, data)
	copy(ciphertext[aes.BlockSize:], data)
	ivg, _ := generateIV(ciphertext, encParams.randInt)
	for i := 0; i < aes.BlockSize; i++ {
		iv[i] = ivg[i]
	}

	// Encode the ciphertext in Base64 before returning
	encodedCiphertext, err := xxtea.EncryptBase64(ciphertext, key, true, 0)
	if err != nil {
		return "", err
	}

	return encodedCiphertext, nil
}

func DecryptXXTEA(key []byte, encodedCiphertext string) ([]byte, error) {
	decB64, err := xxtea.DecryptBase64(encodedCiphertext, key, true, 0)
	if err != nil {
		return nil, err
	}

	return decB64[aes.BlockSize:], nil
}
