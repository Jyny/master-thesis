package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func randStr(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(fmt.Errorf("aes.rand_str: %w", err))
	}

	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}

	return string(bytes)
}

func CFBEncrypt(key, plainText []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(fmt.Errorf("aes.encrypt: %w", err))
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(fmt.Errorf("aes.encrypt: %w", err))
	}

	cfbStream := cipher.NewCFBEncrypter(block, iv)
	cfbStream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return ciphertext
}

func CFBDecrypt(key, cipherText []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(fmt.Errorf("aes.decrypt: %w", err))
	}

	if len(cipherText) < aes.BlockSize {
		panic(fmt.Errorf("aes.decrypt: %w", errors.New("ciphertext too short")))
	}

	iv := cipherText[:aes.BlockSize]
	plainText := make([]byte, len(cipherText[aes.BlockSize:]))

	cfbStream := cipher.NewCFBDecrypter(block, iv)
	cfbStream.XORKeyStream(plainText, cipherText[aes.BlockSize:])

	return plainText
}
