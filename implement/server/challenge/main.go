package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"server/pkg/rsa"
)

func main() {
	if len(os.Args) < 2 {
	} else {
		cipher, err := base64.StdEncoding.DecodeString(os.Args[2])
		if err != nil {
			panic(err)
		}

		plainText, err := rsa.Decrypt(cipher, os.Args[1])
		if err != nil {
			panic(err)
		}

		hash := sha256.Sum256(plainText)
		sign, err := rsa.Sign(os.Args[1], hash[:])
		if err != nil {
			panic(err)
		}

		data, err := json.MarshalIndent(struct {
			Solve string `json:"solve"`
			Sign  string `json:"sign"`
		}{
			Solve: base64.StdEncoding.EncodeToString(plainText),
			Sign:  base64.StdEncoding.EncodeToString(sign),
		}, "", "    ")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(data))
	}
}
