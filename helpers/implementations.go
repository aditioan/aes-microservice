package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// EncryptServiceInstance is the implementation of interface for micro service
type EncryptServiceInstance struct {}

// Implement AES encryption standard (Rijndael Algorithm)
/* Initialization vector for the AES algorithm
More details visit this link
https://en.wikipedia.org/wiki/Advanced_Encryption_Standard */
var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66,32, 14, 05}

var keyEmpty = errors.New("secret Key should not be empty")
var textEmpty = errors.New("secret Text / Message should not be empty")

// Encrypt encrypt string with the given key
func (EncryptServiceInstance) Encrypt(_ context.Context, key, text string) (string, error) {
	if key == "" {
		return "", keyEmpty
	}
	if text == "" {
		return "", textEmpty
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plainttext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	ciphertext := make([]byte, len(plainttext))
	cfb.XORKeyStream(ciphertext, plainttext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the encrypted string to original
func (EncryptServiceInstance) Decrypt(_ context.Context, key, text string) (string, error) {
	if key == "" {
		return "", keyEmpty
	}
	if text == "" {
		return "", textEmpty
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}