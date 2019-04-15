package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

// https://play.golang.org/p/_9zQJ0aWaG

const (
	secret = "JHjh()z&)/hlLZ(jn(jnjnJHJ68JoUu7" // 32 Bytes
)

var (
	v_encrypt *string
	v_decrypt *string
)

func init() {
	v_encrypt = flag.String("encrypt", "", "text to encrypt")
	v_decrypt = flag.String("decrypt", "", "text to decrypt")
}

func main() {
	flag.Parse();
	if len(os.Args) == 1 {
		flag.Usage();
		os.Exit(0)
	}

	if *v_encrypt != "" {
		fmt.Printf("%s\n", "$$"+encrypt(secret, *v_encrypt))
		os.Exit(0)
	}

	if *v_decrypt != "" {
		if !strings.HasPrefix(*v_decrypt, "$$") {
			panic(fmt.Errorf("decrypt phrase must start with a $$"))
		}

		fmt.Printf("%s\n", decrypt(secret, (*v_decrypt)[2:]))
		os.Exit(0)
	}
}

func encrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(key[:16]))
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext)
}

func decrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext, err := hex.DecodeString(text)
	if err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, []byte(key[:16]))
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)

	return string(plaintext)
}
