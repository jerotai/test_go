package jellyfish

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"Stingray/api/jellyfish"
	"net/http"
	"Stingray/helper"
)

var (
	loginJson = `{"account":"CQ1_admin","password":"1234567890","hall_code":"CQ1"}`
	chekcLoginJson = `{"account":"CQ1_admin","password":"1234567890","hall_code":"CQ1"}`
)

/**
 *測試 Login
 */
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	
	r.POST("/login", ex.Login)
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

