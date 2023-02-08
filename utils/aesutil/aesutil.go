package aesutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type AESUtil interface {
	AesEncrypt(data []byte, nonce string) (string, error)
	AesDecrypt(data []byte) (string, error)
	AesVerify(payloads, verifyData string) error
	AesEncryptCBC(data []byte) (string, error)
	AesDecryptCBC(data []byte) (string, error)
	AesVerifyCBC(payloads, verifyData string) error
}
type aesUtil struct {
	key string
	iv  string
}

func NewAesUtil(key string, iv string) AESUtil {
	return &aesUtil{key: key, iv: iv}
}

func (u *aesUtil) AesVerify(payloads, verifyData string) error {
	payloadsByte, err := hex.DecodeString(payloads)
	if err != nil {
		return err
	}
	resp, err := u.AesDecrypt(payloadsByte)
	if err != nil {
		return err
	}
	if resp != verifyData {
		return errors.New("aes verify fail")
	}
	return nil
}

func (u *aesUtil) AesEncrypt(data []byte, nonce string) (string, error) {
	block, err := aes.NewCipher([]byte(u.key))
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(nonce) > aesGCM.NonceSize() {
		return "", fmt.Errorf("nonce is more than %v character", aesGCM.NonceSize())
	}
	cipherText := aesGCM.Seal(nil, []byte(nonce), data, nil)
	return hex.EncodeToString(cipherText), nil
}

func (u *aesUtil) AesDecrypt(data []byte) (string, error) {
	block, err := aes.NewCipher([]byte(u.key))
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	var (
		nonce      = data[:aesGCM.NonceSize()]
		cipherText = data[aesGCM.NonceSize():]
	)
	respData, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", respData), nil
}

func (u *aesUtil) AesVerifyCBC(payloads, verifyData string) error {
	payloadsByte, err := base64.StdEncoding.DecodeString(payloads)
	if err != nil {
		return err
	}
	resp, err := u.AesDecryptCBC(payloadsByte)
	if err != nil {
		return err
	}
	if resp != verifyData {
		return errors.New("aes verify fail")
	}
	return nil
}

func (u *aesUtil) AesEncryptCBC(data []byte) (string, error) {
	block, err := aes.NewCipher([]byte(u.key))
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", errors.New("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(u.iv))
	data = PKCS5Padding(data, block.BlockSize())
	crypted := make([]byte, len(data))
	ecb.CryptBlocks(crypted, data)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (u *aesUtil) AesDecryptCBC(data []byte) (string, error) {
	block, err := aes.NewCipher([]byte(u.key))
	if err != nil {
		fmt.Println("key error", err)
	}
	if len(data) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(u.iv))
	decrypted := make([]byte, len(data))
	ecb.CryptBlocks(decrypted, data)

	return fmt.Sprintf("%s", PKCS5Trimming(decrypted)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
