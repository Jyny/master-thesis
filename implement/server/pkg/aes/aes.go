package aes

import (
	"fmt"
	"io/ioutil"
)

var (
	keyLen = 32
)

func CreateKey() []byte {
	newkey := []byte(randStr(keyLen))

	return newkey
}

func EncryptFile(key []byte, inputfile string, outputfile string) error {
	b, err := ioutil.ReadFile(inputfile)
	if err != nil {
		return fmt.Errorf("aes.EncryptFile: %w", err)
	}

	ciphertext := CFBEncrypt(key, b)
	err = ioutil.WriteFile(outputfile, ciphertext, 0644)
	if err != nil {
		return fmt.Errorf("aes.EncryptFile: %w", err)
	}

	return nil
}

func DecryptFile(key []byte, inputfile string, outputfile string) error {
	cipher, err := ioutil.ReadFile(inputfile)
	if err != nil {
		return fmt.Errorf("aes.DecryptFile: %w", err)
	}

	plain := CFBDecrypt(key, cipher)
	err = ioutil.WriteFile(outputfile, plain, 0644)
	if err != nil {
		return fmt.Errorf("aes.DecryptFile: %w", err)
	}

	return nil
}
