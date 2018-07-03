package jellyfish

import (
	"testing"
)


func TestAuthGroupCreate(t *testing.T) {
	/*gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("jellyfish_service")
	ex.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	ex.SendCurl.ApiToken = ApiToken
	apiPath := "/authGroup"
	
	r.POST(apiPath, ex.HallSendPost)
	req, err := http.NewRequest(http.MethodPost, apiPath, strings.NewReader(createAuthGroupJson))
	
	
	if err != nil {
		t.Fatalf("TestShareholderList NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ex.SendCurl.ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestShareholderList ServeHTTP Error")
	}
	*/
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}