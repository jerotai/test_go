package whale

import (
	"testing"
	"Stingray/api/whale"
	"Stingray/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"strings"
)

func TestCraeteBanner(t *testing.T) {
	ex := whale.New()
	ex.SendCurl = helper.NewCurlBase()
	ex.SendCurl.ApiToken = ApiToken
	apiPath := "/banner"
	
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	r.POST(apiPath, ex.HallSendPost)
	req, err := http.NewRequest(
		http.MethodPost,
		apiPath ,
		strings.NewReader(createBannerJson))
	
	if err != nil {
		t.Fatalf("TestCraeteBanner NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		fmt.Println(w)
		fmt.Println(req)
		t.Fatalf("TestCraeteBanner ServeHTTP Error")
	}
}

func TestMenuList(t *testing.T) {
	
	/*ex := whale.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("WhaleService")
	fmt.Println(apiConf)
	
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	ex.SendCurl.ApiToken = ApiToken
	apiPath := "/menu/list"
	apiParams := "" //
	fmt.Println(apiPath, apiParams)
	
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	r.GET(apiPath, ex.WhaleSendGet)
	req, err := http.NewRequest(http.MethodGet, apiPath  + apiParams, nil)
	
	if err != nil {
		t.Fatalf("TestMenuList NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		fmt.Println(w)
		fmt.Println(req)
		t.Fatalf("TestMenuList ServeHTTP Error")
	}*/
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}