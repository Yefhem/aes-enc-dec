package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	say, err := cowsay.Say(
		"...than it is to decrypt it.",
		cowsay.Type("bud-frogs"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
	fmt.Println("Decryption")
	fmt.Println()

	SecretKey := []byte("")
	ciphertext, err := ioutil.ReadFile("../encrypt/data.data")

	if err != nil {
		fmt.Println(err)
	}
	c, err := aes.NewCipher(SecretKey)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))
}
