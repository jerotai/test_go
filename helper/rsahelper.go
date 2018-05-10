package helper

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"errors"
)

// 加密
func RsaEncrypt(origData string) ([]byte, error) {
	block, _ := pem.Decode([]byte(RsaPublicKey()))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(RsaPrivateKey()))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}