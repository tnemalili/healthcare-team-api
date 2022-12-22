package core

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
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

func VALIDATENumber(mobile string) error {
	startsWithZero := strings.HasPrefix(mobile, "0")
	numberLength, err := strconv.Atoi(os.Getenv("NUMBER_LENGTH"))
	if err != nil {
		return errors.New("invalid number length")
	}
	validLength := len(mobile) == numberLength
	if !startsWithZero || !validLength {
		return errors.New("invalid mobile number")
	}
	return nil
}

func PrependCode(mobile string) string {
	m := strings.TrimPrefix(mobile, "0")
	return fmt.Sprintf("%v%v", os.Getenv("COUNTRY_CODE"), m)
}

func GenerateOTP() (string, error) {

	now := time.Now()
	secret := os.Getenv("OTP_SECRET")
	otpLength, err := strconv.Atoi(os.Getenv("OTP_LENGTH"))
	if err != nil {
		log.Error("Problem getting OTP Length")
		return "0000", err
	}
	digits := otp.Digits(otpLength)
	counter := now.Add(100)
	opts := totp.ValidateOpts{Digits: digits}
	code, err := totp.GenerateCodeCustom(secret, counter, opts)
	if err != nil {
		log.Error("Problem Generating OTP: ", err.Error())
		return "", err
	}
	log.Warnf("[GenerateOTP] OTP: %v", code)

	return code, nil
}

func HashIdentity(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(b), err
}

func CheckIdentityHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
