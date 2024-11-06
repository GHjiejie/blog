package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var (
	// AesKey AES加密的key
	dbAESSk = []byte("blogblogconsole1")
)

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(plaintext []byte) ([]byte, error) {
	plaintextLen := len(plaintext)
	if plaintextLen == 0 {
		return nil, fmt.Errorf("plaintext length(%d) is invalid", plaintextLen)
	}

	paddingLen := int(plaintext[plaintextLen-1])
	if paddingLen > plaintextLen {
		return nil, fmt.Errorf("paddingLen length(%d) should be less than plaintext length(%d), while not",
			paddingLen, plaintextLen)
	}
	return plaintext[:plaintextLen-paddingLen], nil
}

// 将用户的密码进行加密之后，存入数据库
func AESEncrypt(plaintext []byte) (string, error) {
	if len(plaintext) == 0 {
		return "", fmt.Errorf("empty plaintext")
	}

	block, err := aes.NewCipher(dbAESSk)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = pkcs7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, dbAESSk[:blockSize])
	ciphertextBytes := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertextBytes, plaintext)
	return fmt.Sprintf("%x", ciphertextBytes), nil
}

// 从数据库中获取用户密码进行解密
func AESDecrypt(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", fmt.Errorf("empty ciphertext")
	}

	block, err := aes.NewCipher(dbAESSk)
	if err != nil {
		return "", err
	}

	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(ciphertextBytes)%blockSize != 0 {
		return "", fmt.Errorf("invalid crypted length(%d), not a multiple of the block size(%d)",
			len(ciphertextBytes), blockSize)
	}

	blockMode := cipher.NewCBCDecrypter(block, dbAESSk[:blockSize])
	plaintext := make([]byte, len(ciphertextBytes))
	blockMode.CryptBlocks(plaintext, ciphertextBytes)
	plaintext, err = pkcs7UnPadding(plaintext)
	return string(plaintext), err
}
