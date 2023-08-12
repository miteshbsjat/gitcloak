package encrypt

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/yang3yen/xxtea-go/xxtea"
)

// var ENCRYPTION_ALGORITHMS []string = []string{"aes", "chacha", "xxtea"}

func Encrypt(line []byte, line_num int, encryptionAlgo string, key []byte) (string, error) {
	var encs []byte = nil
	// salt := []byte(strconv.Itoa(line_num % len(key)))
	// saltLine := append(line, salt[len(salt)-1])
	saltLine := line
	var err error = nil
	switch encryptionAlgo {
	case "xxtea":
		encs, err = xxtea.Encrypt(saltLine, key, true, 0)
	default:
		encs, err = xxtea.Encrypt(saltLine, key, true, 0)
	}
	if err != nil {
		fmt.Println(err)
	}
	encb := b64.StdEncoding.EncodeToString(encs)
	return encb, nil
}
func Decrypt(line string, encryptionAlgo string, key []byte) ([]byte, error) {
	var encs []byte = nil
	encb, _ := b64.StdEncoding.DecodeString(line)
	switch encryptionAlgo {
	case "xxtea":
		encs, _ = xxtea.Decrypt(encb, key, true, 0)
	default:
		encs, _ = xxtea.Decrypt(encb, key, true, 0)
	}
	// return encs[0 : len(encs)-1], nil
	return encs, nil
}
