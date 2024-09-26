package aesUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/mooncake9527/x/xerrors/xerror"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, xerror.New(err.Error())
	}
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, make([]byte, block.BlockSize()))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

func Decrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, xerror.New(err.Error())
	}
	decrypted := make([]byte, len(encrypted))
	stream := cipher.NewCTR(block, make([]byte, block.BlockSize()))
	stream.XORKeyStream(decrypted, encrypted)
	return decrypted, nil
}
