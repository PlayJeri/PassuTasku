package manager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

var str = "passuTaskufjdksL"
var bytes = []byte(str)

func Encrypt(plaintext []byte) []byte {
	copy(bytes[:], str)
	block, err := aes.NewCipher(bytes)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphered := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphered
}

func Decrypt(ciphered []byte) []byte {
	block, err := aes.NewCipher(bytes)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	if len(ciphered) < gcm.NonceSize() {
		panic("cipher text too short")
	}

	nonce, ciphered := ciphered[:gcm.NonceSize()], ciphered[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphered, nil)
	if err != nil {
		panic(err)
	}

	return plaintext
}
