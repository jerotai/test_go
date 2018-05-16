package helper

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// 加密
func RsaEncrypt(origData string) (string, error) {
	block, _ := pem.Decode([]byte(RsaPublicKey()))
	if block == nil {
		errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	rsaEnode, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(rsaEnode), nil
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
	
	signed, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}
	
	return rsa.DecryptPKCS1v15(rand.Reader, priv, signed)
}