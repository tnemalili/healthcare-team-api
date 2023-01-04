package core

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var bytesArray = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func BOD(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GenerateCode() string {
	return strings.ToUpper(guuid.NewString()[:7])
}

func GenerateTransactionId() string {
	return strings.ToUpper(guuid.NewString())
}
func Now() time.Time {

	return time.Now()
}

func EncryptAES(plaintext string) (string, error) {

	key := os.Getenv("CIPHER_KEY")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	plainText := []byte(plaintext)
	cfb := cipher.NewCFBEncrypter(block, bytesArray)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func DecryptAES(ct string) (string, error) {

	key := os.Getenv("CIPHER_KEY")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	cipherText := Decode(ct)
	cfb := cipher.NewCFBDecrypter(block, bytesArray)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {

	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Error(err)
	}
	return data
}

func ElapsedTime(startTime time.Time, endTime time.Time) int {

	return int(endTime.Sub(startTime).Seconds())
}

func ISReservedUser(username string) bool {
	return username == os.Getenv("RESERVED_USERNAME")
}

func ISReservedPassword(username string) bool {
	return username == os.Getenv("RESERVED_PASSWORD")
}

func Verified(inputValue interface{}, expectedValue interface{}) bool {

	return inputValue == expectedValue
}
