package jellyfish

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"Stingray/api/jellyfish"
	"net/http"
	"Stingray/helper"
	"fmt"
)

var (
	hallJson = `{"account":"testaccount","password":"1234567890","name":"testname","new_hall_code":"CQ111"}`
	chekcHallJson = `{"account":"test","password":"1234567890","hall_code":"CQ1"}`
	hallPutJson = `{"id":"2","name":"testname"}`
	hallenabledJson = `{"id":"14"}`
	subListJson = `{"page":"1","count":"10"}`
	ApiToken = "$2y$10$ItiEb3PhAwRwERrY9rGN\\/eUrMdBb\\/dk025ndmvgR.1B3phHWZRcLK"
)

/**
 *測試 Hall Post
 */
func TestHallPost(t *testing.T) {
	rsaData, _ := helper.RsaEncrypt(hallJson)
	
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	ex.SendCurl.ApiToken = ApiToken
	
	r.POST("/hall", ex.Hall)
	
	req, err := http.NewRequest(http.MethodPost, "/hall", strings.NewReader(string(rsaData)))
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Api-Token", ex.SendCurl.ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("TestHall ServeHTTP Error : %v", req)
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}

/**
 *測試 Hall Put
 */
func TestHallPut(t *testing.T) {
	rsaData, _ := helper.RsaEncrypt(hallPutJson)
	
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	ex.SendCurl.ApiToken = ApiToken
	
	r.PUT("/hall", ex.UpdataHall)
	req, err := http.NewRequest(http.MethodPut, "/hall", strings.NewReader(string(rsaData)))
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Api-Token", ex.SendCurl.ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	fmt.Println(w.Code)
	if w.Code != http.StatusOK {
		t.Fatalf("TestHallPut ServeHTTP Error : %v", req)
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}

/**
 *測試 EnabledHall
 */
func TestEnabledHall(t *testing.T) {
	rsaData, _ := helper.RsaEncrypt(hallenabledJson)
	
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	ex.SendCurl.ApiToken = ApiToken
	apiPath := "/hall/enabled"
	
	r.PUT(apiPath, ex.EnabledHall)
	req, err := http.NewRequest(http.MethodPut, apiPath, strings.NewReader(string(rsaData)))
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ex.SendCurl.ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	//req.RequestURI = "CQ1.admin.com"
	
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("TestHallPut ServeHTTP Error : %v", req)
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}

func TestSubList(t *testing.T) {
	//rsaData, _ := helper.RsaEncrypt(hallenabledJson)
	
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("JellyFishService")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	ex.SendCurl.ApiToken = ApiToken
	apiPath := "/hall/subList"
	
	r.GET(apiPath + "/:Page/:Count", ex.SubList)
	req, err := http.NewRequest(http.MethodGet, apiPath + "/1/10", nil)
	
	if err != nil {
		t.Fatalf("TestLogin NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ex.SendCurl.ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestSubList ServeHTTP Error")
	}
	
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}