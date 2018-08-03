package whale

import (
	"testing"
	"github.com/gin-gonic/gin"
	"os"
	"Stingray/helper"
	"fmt"
	"Stingray/core/model/redis/rsakey"
	"net/http/httptest"
	"Stingray/api/whale"
	"Stingray/core/apicurl"
	"strings"
	"net/http"
)

func TestFocus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("ROUTER_RSA_OPEN","Y")
	
	//取得指定 token 的 rsa private key
	token := ApiToken
	rsaRedisService := rsakey.New()
	rsaRedisService.GetTokenRsaPublicKey(token)
	rsaKey := rsaRedisService.GetTokenRsaPublicKey(token)
	
	//input rsa encode
	rsaencode, err := helper.RsaEncryptByKey(rsaKey, focusJson)
	
	if err != nil {
		t.Fatalf("TestFocus RsaEncryptByKey Error : %v", err)
	}
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	apiConf := helper.ApiSetting("whale_service")
	router := whale.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	var apiCurlSend, _ = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	r.POST("/bookmark/focus", apiCurlSend.HallSendPost)
	req, err := http.NewRequest(http.MethodPost, "/bookmark/focus", strings.NewReader(rsaencode))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ApiToken)
	req.Header.Set("Origin", "http://dodo.cqcp.corp")
	
	if err != nil {
		t.Fatalf("TestFocus NewRequest Error : %v", err)
	}
	
	//req.Header.Set("Content-Type", "multipart/form-data")
	
	r.ServeHTTP(w, req)
	fmt.Println(w)
	if w.Code != http.StatusOK {
		t.Fatalf("TestLogin ServeHTTP Error : %v", err)
	}
}