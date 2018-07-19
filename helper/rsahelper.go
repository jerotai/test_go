package helper

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

/**
 * 加密 by key
 */
func RsaEncryptByKey(key string, origData string) (string, error) {
	block, _ := pem.Decode([]byte(key))
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

/**
 * 解密 by key
 */
func RsaDecryptByKey(key string, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(key))
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

/**
 * Create Login Rsa Key
 */
func CreateRsaKey() (string, string){
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		HelperLog.ErrorLog("[rsahelper CreateRsaKey] Private key cannot be created. " + err.Error())
		return "", ""
	}
	
	// Generate a pem block with the private key
	prikeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	// Generate a pem block with the public key
	PubASN1, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	pubkeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})
	
	if err != nil {
		HelperLog.ErrorLog("[rsahelper CreateRsaKey] pubkeyPem cannot be created. " + err.Error())
		return "", ""
	}
	
	return string(pubkeyPem), string(prikeyPem)
}