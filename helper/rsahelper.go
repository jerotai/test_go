package helper

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"bytes"
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
	publicKey := pubInterface.(*rsa.PublicKey)
	
	partLen := publicKey.N.BitLen() / 8 - 11
	chunks := Split([]byte(origData), partLen)
	
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
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
 * new rsa decode
 */
func RsaDecryptByKeyBySplit(key string, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	
	partLen := priv.PublicKey.N.BitLen() / 8
	
	signed, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}
	
	chunks := Split([]byte(signed), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priv, chunk)
		if err != nil {
			return nil, err
		}
		
		buffer.Write(decrypted)
	}
	
	return []byte(buffer.String()), nil
}

/**
 * Create Login Rsa Key
 */
func CreateRsaKey() (string, string){
	key, err := rsa.GenerateKey(rand.Reader, 4096) //use RsaDecryptByKeyBySplit 4096 change 1024
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

func Split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}