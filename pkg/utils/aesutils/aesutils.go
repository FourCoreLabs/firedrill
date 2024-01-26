package aesutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// aesEncryptData encrypts data using 256-bit AES-GCM. Output: nonce+cipherdata+tag
func AESEncryptData(data []byte, key []byte) (encryptedtext []byte, err error) {
	cipherblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcmpack, err := cipher.NewGCM(cipherblock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcmpack.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcmpack.Seal(nonce, nonce, data, nil), nil
}

// aesEncryptionKey returns random AES Encrpytion Key
func AESEncryptionKey() []byte {
	ekey := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, ekey)
	if err != nil {
		panic(fmt.Sprintf("Failed to seed key: %v", err))
	}
	return ekey
}
