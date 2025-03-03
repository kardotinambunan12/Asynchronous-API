package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"sync"
)

var key = []byte("1234567890kardo")

func ConvertToSHA1(value string) string {
	sha := sha1.New()
	sha.Write([]byte(value))
	encrypted := sha.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)
	return encryptedString
}

func GenerateSignature(data, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(hashed)
	return signature
}

func GenerateCounter() int {
	var (
		counter int
		mu      sync.Mutex
	)

	mu.Lock()
	defer mu.Unlock()

	counter++

	return counter
}

func GenerateLastCounter(counter int) int {
	var (
		mu sync.Mutex
	)

	mu.Lock()
	defer mu.Unlock()

	counter++

	return counter
}
