package helper

import (
	"testing"
	"Stingray/helper"
	"fmt"
	"crypto/rsa"
	"encoding/base64"
	"bytes"
	"crypto/rand"
	"encoding/pem"
	"crypto/x509"
	"github.com/magiconair/properties/assert"
)

var pubKey = `-----BEGIN RSA PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKmYOc+S+AdJN+RlAJI49JZchCsl+8jP
7U3KY4SmAIc4B+csBzo9ojSU6U7OhT5N7+H/nvnJNz12ZHhlsmnWRKcCAwEAAQ==
-----END RSA PUBLIC KEY-----`

var priKey = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBAKmYOc+S+AdJN+RlAJI49JZchCsl+8jP7U3KY4SmAIc4B+csBzo9
ojSU6U7OhT5N7+H/nvnJNz12ZHhlsmnWRKcCAwEAAQJBAJ5fGR1OEXA8X4VZDH9/
0GfZXVMt1UJhaSkLpoNowXWCx/3N0+S5ZQmbFTHKJrSVzMM03GMzczfVSQaCA1EF
LQECIQDJf/35a+ns6ajMfBz6rFHeK1vjLcmniXEWBkJgr58PvwIhANd3FZ455EyJ
oil8SsVnxINi+QBubvhDvymE5gn3wQUZAiEAimGsOi8yZ9HQax6RiFXvo3SDlnFL
nBN35nhF/cCerGECIDDwxaJI+AhCguj1aG5qYYKOaoykiOQvTy55F19QPTGRAiBn
mLbGXHzh00L+aLHYr+NlQLsiPPX6lNHKR6iJUzkHoQ==
-----END RSA PRIVATE KEY-----`

var testData = `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`

// 測試 RSA By Key
// RsaEncryptByKey
// RsaDecryptByKey
func TestRsaByKey(t *testing.T) {
	data := "test"
	testEnRsaString, err := helper.RsaEncryptByKey(pubKey, data)
	
	if err != nil {
		t.Error("helper.RsaEncryptByKey Error")
	}
	
	testDeRsaString, err := helper.RsaDecryptByKeyBySplit(priKey, []byte(testEnRsaString))
	
	if err != nil {
		t.Error("helper.RsaEncryptByKey Error")
	}
	
	assert.Equal(t, string(testDeRsaString), data)
}

func TestCreateRsaKey(t *testing.T) {

	pub, pri := helper.CreateRsaKey()
	fmt.Println(pub)
	if pub == "" {
		t.Error("CreateRsaKey Pubkey Error")
	}
	fmt.Println(pri)
	if pri == "" {
		t.Error("CreateRsaKey Prikey Error")
	}
}

func TestSplitRsa(t *testing.T) {
	rsaEncode, err := publicEncrypt(testData)
	if (err != nil) {
		fmt.Println(err)
	}
	
	fmt.Println(rsaEncode)
	
	rsaDecode, err := privateDecrypt(rsaEncode)
	
	fmt.Println(rsaDecode)
}

func rsaPublic() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		//errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := pubInterface.(*rsa.PublicKey)
	return publicKey
}

func rsaPrivate() *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(priKey))
	
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	
	return priv
}

func publicEncrypt(data string) (string, error) {
	publicKey := rsaPublic()
	
	partLen := publicKey.N.BitLen() / 8 - 11
	chunks := helper.Split([]byte(data), partLen)
	
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}

	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}

func privateDecrypt(encrypted string) (string, error) {
	
	publicKey := rsaPublic()
	private := rsaPrivate()
	
	partLen := publicKey.N.BitLen() / 8
	raw, err := base64.RawURLEncoding.DecodeString(encrypted)
	chunks := helper.Split([]byte(raw), partLen)
	
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, private, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}
	
	return buffer.String(), err
}