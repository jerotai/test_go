package jellyfish

import (
	"testing"
)


func TestShareholderList(t *testing.T) {
	/*gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := jellyfish.New()
	ex.SendCurl = helper.NewCurlBase()
	apiConf := helper.ApiSetting("jellyfish_service")
	ex.SendCurl.Url = fmt.Sprintf("%s:%d/", apiConf.Host, apiConf.Port)
	ex.SendCurl.ApiToken = ApiToken
	apiPath := "/shareholder/list"
	apiParams := "?Status=asd&Start_Time=2018-01-01&End_Time=12345-1235-22" //
	
	r.GET(apiPath + "/:Site_Code/:Page/:Count/", ex.HallSendGet)
	req, err := http.NewRequest(http.MethodGet, apiPath + "/AA001/1/10/" + apiParams, nil)
	
	if err != nil {
		t.Fatalf("TestShareholderList NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", ApiToken)
	req.Header.Set("Origin", "http://CQ1.admin.com:8080")
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		fmt.Println(w)
		fmt.Println(req)
		t.Fatalf("TestShareholderList ServeHTTP Error")
	}
	*/
	//assert.Equal(t,  w.Body.String(), string(chekcLoginJson))
}