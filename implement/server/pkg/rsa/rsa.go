package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

var (
	bitSize = 1024
)

func GenerateKey() (pk, sk string, err error) {
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return "", "", fmt.Errorf("rsa.Decrypt: %w", err)
	}

	pk = base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&key.PublicKey))
	sk = base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(key))

	return pk, sk, nil
}

func Decrypt(ciphertext []byte, privateKey string) ([]byte, error) {
	rsaPrivB, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("rsa.DecryptStr: %w", err)
	}

	rsaPriv, err := x509.ParsePKCS1PrivateKey(rsaPrivB)
	if err != nil {
		return nil, fmt.Errorf("rsa.Decrypt: %w", err)
	}

	return rsa.DecryptPKCS1v15(rand.Reader, rsaPriv, ciphertext)
}

func Encrypt(plainText []byte, publicKey string) ([]byte, error) {
	rsaPubB, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("rsa.EncryptStr: %w", err)
	}

	rsaPub, err := x509.ParsePKCS1PublicKey(rsaPubB)
	if err != nil {
		return nil, fmt.Errorf("rsa.Encrypt: %w", err)
	}

	return rsa.EncryptPKCS1v15(rand.Reader, rsaPub, plainText)
}
