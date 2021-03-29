// crypto/cipher: For block cipher implementation.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func encrypt(data []byte, key string) (resp []byte, err error) {
	// Cifra la clave
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	// con la clave crifrada crea el gcm
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	var dst []byte
	// Crea un array nonce del tamaño del gcm
	nonce := make([]byte, gcm.NonceSize())
	dst = nonce
	if _, err := rand.Read(nonce); err != nil {
		return resp, err
	}
	return gcm.Seal(dst, nonce, data, []byte("test")), nil
}

func decrypt(data []byte, key string) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}

	ciphertext := data[gcm.NonceSize():]

	nonce := data[:gcm.NonceSize()]
	resp, err = gcm.Open(nil, nonce, ciphertext, []byte("test"))
	if err != nil {
		return resp, fmt.Errorf("error decrypting data: %v", err)
	}
	return resp, nil
}

func main() {
	// La clave debe tener 16 caracteres
	const key_enc = "mysecurepassword"
	const key_dec = "mysecurepassword"

	clear_text := "Hello World!"
	fmt.Println("Clear text: ", clear_text)
	encrypted, err := encrypt([]byte(clear_text), key_enc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Encrypted Text: %v \n", string(encrypted))
	decrypted, err := decrypt(encrypted, key_dec)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decrypted Text: ", string(decrypted))
}
