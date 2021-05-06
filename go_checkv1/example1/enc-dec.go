package example1

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type Blob struct {
	Data       []byte
	CipherText []byte
}

func (b *Blob) Encrypt(key []byte) []byte {
	//initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("failed in creating block cipher")
	}
	// create  byte slice the will hold the encrypted message
	cipherText := make([]byte, aes.BlockSize+len(b.Data))

	// generate the initialization Vector (VI) nonce
	//which will be store at the beginning of the byte slice. the VI is the same length as the AES blockSize
	vi := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, vi)
	if err != nil {
		log.Fatal("failed in filling the vi")
	}
	//choose the block cipher mode of operation.
	//using the cipher feedback (CFB) mode here. CBCEncrypter is  also available
	cfb := cipher.NewCFBEncrypter(block, vi)
	//generate the encrypted message and store it in the remaining bytes after vi.
	cfb.XORKeyStream(cipherText[aes.BlockSize:], b.Data)
	return cipherText
}

func (b *Blob) Decryption(key []byte) []byte {
	cipherText := b.CipherText
	//initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("failed in creating a block cipher")
	}
	//separate the vi nonce from the encrypted message
	vi := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	//decrypt the message using CFB block mode
	cfb := cipher.NewCFBDecrypter(block, vi)
	cfb.XORKeyStream(cipherText, cipherText)
	return cipherText
}
