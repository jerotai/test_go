package jellyfish

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"Stingray/api/jellyfish"
	"net/http"
	"Stingray/helper"
	"Stingray/core/model/redis/rsakey"
	"fmt"
)


/**
 *測試 Login
 */
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	//取得指定 token 的 rsa private key
	token := ApiToken
	rsaRedisService := rsakey.RsaKeyService()
	rsaRedisService.Init()
	rsaKey := rsaRedisService.GetTokenRsaPrivateKey(token)
	fmt.Println(rsaKey)
	//input rsa encode
	rsadecode, _ := helper.RsaDecryptByKey(rsaKey, []byte("b2W85bT1Gd3TlgVQslbs6gBnopXbg7EBhi/LTx310nRhsI8BL6/9YDUrDuGpIjNH5Hzc7EZBTK7yLRhDmZbyzte9Oe9br5aU6vV4B3Bow18dJ7VC3RhVzrVuT45pn1tNG1S56nB+NbIA3chnjLC3EHx8bMxfkpyYFLQMR/WUqGhc7qpiTLbt7B2T7Ktgw0bmwKVoyJwtjVzlM/7Srx3OCbHqyYr4uv9tK57F+TYiPPzZY7zb1CK23xwQU6V6qhAlk9y6R1xeBrZy2HfgpFGuC8/saAVXVUtt/ds8eZ7lagIr4ojWr57B3DF7laMPAalKJlhdvk8nhY6Jzaz0i+cX7g=="))
	
	fmt.Println(rsadecode)
	
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	
	r.POST("/login", ex.HallSendPost)
	req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(loginJson))
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestLogin ServeHTTP Error : %v", err)
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}

