package apiparse

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"routes/helper"
	"routes/core/dto"
	"fmt"
)

/**
 * 回傳格式彙整
 */
func ApiResponse(req interface{}) interface{} {
	statusReq := dto.ApiStatus{}
	b, _ := json.Marshal(req)
	
	statusReq.Code = statusCodeParse("1")
	statusReq.Result = string(b[:])
	return statusReq
}

/*func buildQuery(c echo.Context) {
	q := c.Request().URL.Query()
	for k, v := range c.paramQuery {
		q.Set(k, v)
	}
	c.RawQuery = q.Encode()
}*/

var m map[string]interface{}

func GetJsonDataParse(c echo.Context, req interface{})  {
	keys := c.Request().ParseForm
	
	fmt.Println(keys)
	/*for _, k := range q {
		req.Set(k, v)
	}*/
	/*if err != nil {
		//後續處理
	}*/
	/*for _, key := range c.ParamNames() {
		fmt.Println(key)
	}*/
	//fmt.Println(req)
}

//取得 post 的參數
func PostJsonDataParse(c echo.Context, req interface{}) error {
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		//後續處理
	}
	
	return err
}

//取得 rsa加密的post 參數
func PostJsonRsaDataParse(c echo.Context, req interface{}) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		//後續處理
	}
	
	rsadecode, _ := helper.RsaDecrypt(body)
	json.Unmarshal([]byte(rsadecode), &req)
	
	return err
}

/**
 * 確認站台 Code
 * url := siteCodeParse(c.Request().Host)
 */
func SiteCodeParse(url string) string {

	return "S1"
}

/**
 api 錯誤碼整理
 */
func statusCodeParse(status string) string {
	statusCode := "1"
	return statusCode
}

// getHost tries its best to return the request host.
/*
func getHost(r *http.Request) string {
	if r.URL.IsAbs() {
		host := r.Host
		// Slice off any port information.
		if i := strings.Index(host, ":"); i != -1 {
			host = host[:i]
		}
		return host
	}
	return r.URL.Host
}
*/

