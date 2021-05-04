package main

import (
	"fmt"
	"os"
	"server/pkg/aes"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(string(aes.CreateKey()))
	} else {
		aes.EncryptFile([]byte(os.Args[1]), os.Args[2], os.Args[1])
	}
}
