package rsa

import (
	"crypto"
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
		return nil, fmt.Errorf("rsa.Decrypt: %w", err)
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
		return nil, fmt.Errorf("rsa.Encrypt: %w", err)
	}

	rsaPub, err := x509.ParsePKCS1PublicKey(rsaPubB)
	if err != nil {
		return nil, fmt.Errorf("rsa.Encrypt: %w", err)
	}

	return rsa.EncryptPKCS1v15(rand.Reader, rsaPub, plainText)
}

func Verify(publicKey string, hash, signature []byte) (bool, error) {
	rsaPubB, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, fmt.Errorf("rsa.Verify: %w", err)
	}

	rsaPub, err := x509.ParsePKCS1PublicKey(rsaPubB)
	if err != nil {
		return false, fmt.Errorf("rsa.Verify: %w", err)
	}

	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, hash, signature)
	if err != nil {
		return false, fmt.Errorf("rsa.Verify: %w", err)
	}

	return true, nil
}

func Sign(publicKey string, hash []byte) ([]byte, error) {
	rsaPrivB, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("rsa.Sign: %w", err)
	}

	rsaPriv, err := x509.ParsePKCS1PrivateKey(rsaPrivB)
	if err != nil {
		return nil, fmt.Errorf("rsa.Sign: %w", err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPriv, crypto.SHA256, hash)
	if err != nil {
		return nil, fmt.Errorf("rsa.Sign: %w", err)
	}

	return signature, nil
}
