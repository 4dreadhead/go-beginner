package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		panic("SECRET_KEY not provided")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	aesGCM := enc.prepare()
	nonce := make([]byte, aesGCM.NonceSize())

	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}

	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	aesGCM := enc.prepare()
	nonceSize := aesGCM.NonceSize()

	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	} 

	return plainText
}

func (enc *Encrypter) prepare() cipher.AEAD {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM
}