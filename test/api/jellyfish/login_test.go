package jellyfish

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"routes/api/jellyfish"
	"net/http"
	"routes/helper"
	"github.com/magiconair/properties/assert"
)

var (
	loginJson = `{"account":"test","password":"1234567890"}`
	chekcLoginJson = `{"account":"test","password":"1234567890","hall_code":"CQ1"}`
)

/**
 *測試 Login
 */
func TestLogin(t *testing.T) {
	rsaData, _ := helper.RsaEncrypt(loginJson)
	
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = &helper.CurlBase{}
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	
	
	r.POST("/login", ex.Login)
	req, err := http.NewRequest("POST", "/login", strings.NewReader(string(rsaData)))
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Host = "CQ1.admin.com"
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestLogin ServeHTTP Error : %v", err)
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}

