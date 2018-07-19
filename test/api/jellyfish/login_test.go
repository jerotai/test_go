package jellyfish

import (
	"testing"
	"Stingray/helper"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"Stingray/api/jellyfish"
	"strings"
	"net/http"
	"Stingray/core/apicurl"
	"fmt"
	"os"
)


/**
 *測試 Login
 */
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("ROUTER_RSA_OPEN","Y")
	//取得指定 token 的 rsa private key
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	apiConf := helper.ApiSetting("jellyfish_service")
	router := jellyfish.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	var apiCurlSend = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	
	r.POST("/login", apiCurlSend.JellyFishLogin)
	req, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(loginJson))
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Origin", "http://cqcp.cqcp.corp")
	
	r.ServeHTTP(w, req)
	fmt.Println(w)
	if w.Code != http.StatusOK {
		t.Fatalf("TestLogin ServeHTTP Error : %v", err)
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}

