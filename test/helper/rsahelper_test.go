package helper

import (
	"testing"
	"Stingray/helper"
	"github.com/magiconair/properties/assert"
)

var pubKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzHPfuhBOuqaQldB9zM3A
j2l5gBMXe4c+lnFARyRrZLMFq1OwnYLclh4XFCmQBEk/uTlxf1jZX9+LX4Q7JCPi
rF4/UctrNj66P+WloTB9yN/jfbu4Owalc2ob9MDFlwKkV1k0IjXe+1jzkYTw+xlj
XvU0rQljNfyUfzTLgKrhkjDSBw8d8k1TBjsnInPPNUqsOn4MBsD2Fejy+4hpbK2S
Tx8s5FFb+1+It8n53al8d6YlcQIB+9qfzxi83uI0ypsE1Vb87B0Niyka3rqKJxN0
9oO1g1OsEXV1zkqXfdlxRJYiBxpyLnw2huVuVlPgRRJqKNGznvJMORKhII2C3EnW
5wIDAQAB
-----END RSA PUBLIC KEY-----`

var priKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAzHPfuhBOuqaQldB9zM3Aj2l5gBMXe4c+lnFARyRrZLMFq1Ow
nYLclh4XFCmQBEk/uTlxf1jZX9+LX4Q7JCPirF4/UctrNj66P+WloTB9yN/jfbu4
Owalc2ob9MDFlwKkV1k0IjXe+1jzkYTw+xljXvU0rQljNfyUfzTLgKrhkjDSBw8d
8k1TBjsnInPPNUqsOn4MBsD2Fejy+4hpbK2STx8s5FFb+1+It8n53al8d6YlcQIB
+9qfzxi83uI0ypsE1Vb87B0Niyka3rqKJxN09oO1g1OsEXV1zkqXfdlxRJYiBxpy
Lnw2huVuVlPgRRJqKNGznvJMORKhII2C3EnW5wIDAQABAoIBAEPD3FG1egLPlnix
hCDPyZU/JnBW45+j8hC3NNDft2DHHYft00pBx49SJeAe7ocdKCviaEYHUvU+CNn9
4ARYiJcNHfukasKuA9mN6loE0owz+RkBkUyvJUOd1epHwrYMwB/bhzYXmPLCRuji
oWT4w0EEERORWvrX3vhSrWjWOrHTODCegdbdBj3syZWxwCs7IDZR+0maq1KNKkQf
FpL+NABXmrbl/Z775AJht16IT8R0RMyincD1xUu8ffx1gugbjuvS0T79TSD4u2PP
6RZ2JDMKJ6NEfx9QP7unG+WZMwvy/y4Fqx8/+zr97hBRGGx/panAur1uMcwpcQsj
0BxZnGECgYEA/cS9pg3Zpu6QBTagctRpGZXY1jPMtMrpsBjI6cI+fXJwUtSff+Ta
sbpAlTnFiij80tFuq1igogjKsnNSqXMDnQQe8qkvSxtYzDcGeeOwIwuGj6hSCtxa
P9EdS0or3zlWEpPQuT2w4s8+04MqhWxBQuThihVO4YnqNqiY1Nfbta8CgYEAzkAe
Mmnr/KtHYzrFdLETmJzvfFohwkfDH4PEKE/v9CPAwMOfzGz1PHVfarft3b6lELWY
64m0iB27CMhXJbvg9QkttBz/S2z8o6ToHs1I782jqn5xcaTuuIwrcN6NP2RYs62X
cddHIzijjT7LcOPk82TEbttYtQWHiFNdOV1XeEkCgYApxM8+XLs9abjU0tf37pRZ
/LsTDrWb8GYkcP/o0er9NO/eAlc6cs24QSLOLRMhmXt7q7MeitjtsqrCUo8BiwfG
OV5qrIQ3RYoJGul6+IoOSqBcVuUiGoDUTjQXJ18vP722ExRg2RGbU2dQ4x8FlPrs
6sVgt/8iylN9qaR2LA1iywKBgC8RWtLOcVhcDhKLGvyiXIJ7bfvG5eVcUx8iBuXc
ODpS3m4fNpyEUQLXFktYZiQwmaLSjm9SIazefAxpC9pEofXJfnCYHXqtCIqf24kL
+BUhrLP/3wu8Q+5Brrg3UayhzcxTVThZlJZpGonlRrAgkeMpn9pcnEXTWTa3L5Mj
xg5hAoGAILHxVl4tMVJi39hiFWbqhqrfoMqhP5qwd/dyTsrV5UTdeJPsKGM4qtpC
jABbXOvsZw4kKZpLbezsxWwrm0LqnevLlKqLBKPh/uRjplXMcC7p9tlb3q3PpzOs
vGiFkiAWaN19MNlew2furIDgwWXnnUFX7VHGrrN6RliqdDmgqQI=
-----END RSA PRIVATE KEY-----`

// 測試 RSA By Key
// RsaEncryptByKey
// RsaDecryptByKey
func TestRsaByKey(t *testing.T) {
	data := "test"
	testEnRsaString, err := helper.RsaEncryptByKey(pubKey, data)
	
	if err != nil {
		t.Error("helper.RsaEncryptByKey Error")
	}
	
	testDeRsaString, err := helper.RsaDecryptByKey(priKey, []byte(testEnRsaString))
	
	assert.Equal(t, string(testDeRsaString), data)
}

func TestCreateRsaKey(t *testing.T) {

	pub, pri := helper.CreateRsaKey()

	if pub == "" {
		t.Error("CreateRsaKey Pubkey Error")
	}
	
	if pri == "" {
		t.Error("CreateRsaKey Prikey Error")
	}
}