package helper

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
	"fmt"
	"errors"
	"bytes"
)

type CurlBase struct {
	cookies []*http.Cookie
	client *http.Client
	ApiToken string
}

//初始化
func NewCurlBase() *CurlBase {
	curlInst := &CurlBase{}
	
	curlInst.client = &http.Client{}
	//為所有重定向的請求增加cookie
		/*curlInst.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) > 0 {
				for _, v := range curlInst.GetCookie() {
					req.AddCookie(v)
				}
			}
			return nil
		}*/
	//})
	
	return curlInst
}

//設置代理地址
func (self *CurlBase) SetProxyUrl(proxyUrl string)  {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	transport := &http.Transport{Proxy:proxy}
	self.client.Transport = transport
}

//設置請求cookie
func (self *CurlBase) AddCookie(cookies []*http.Cookie)  {
	fmt.Println("AddCookie", cookies)
	self.cookies = append(self.cookies, cookies...)
	//self.cookies = cookies
}

//獲取當前所有的cookie
func (self *CurlBase) GetCookie() ([]*http.Cookie) {
	return self.cookies
}

/**
 * CURL GET
 */
func (self *CurlBase) Get(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	getData := self.encodeParams(params)
	apiUrl := requestUrl + "?" + getData
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"NewRequest Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Api-Token", self.ApiToken)
	request.Header.Add("Connection", "close")
	//self.setRequestCookie(request)
	response, err := self.client.Do(request)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"client.Do Failed",
			apiUrl,
			err.Error(),
		)
		
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	defer response.Body.Close()
	
	//保存 cookie
	/*respCks := response.Cookies()
	appendCookie := self.cookies
	self.cookies = append(appendCookie, respCks...)*/
	
	data, err := ioutil.ReadAll(response.Body)
	HelperLog.TraceLog("GET SEND: " + apiUrl)
	HelperLog.TraceLog("CURL GET : " + string(data))
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	//gCurCookies = gCurCookieJar.Cookies(request.URL)
	return data, response.StatusCode, err
}

/**
 * CURL MultiPart
 */
func (self *CurlBase) MultiPartPost(requestUrl string, params *bytes.Buffer, sendContentType string) ([]byte, int, error) {
	
	apiUrl := requestUrl
	request, err := http.NewRequest("POST", apiUrl, params)
	HelperLog.TraceLog("MultiPartPost SEND: " + apiUrl + params.String())
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"NewRequest Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	request.Header.Set("Content-Type", sendContentType)
	request.Header.Add("Api-Token", self.ApiToken)
	//request.Header.Add("Connection", "close")
	//self.setRequestCookie(request)
	response, err := self.client.Do(request)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"client.Do Failed",
			apiUrl,
			err.Error(),
		)
		
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	defer response.Body.Close()
	
	//保存 cookie
	/*respCks := response.Cookies()
	fmt.Println("respCks", respCks)
	appendCookie := self.cookies
	self.cookies = append(appendCookie, respCks...)*/
	
	data, err := ioutil.ReadAll(response.Body)
	
	HelperLog.TraceLog("CURL MultiPartPost : " + string(data))
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	//gCurCookies = gCurCookieJar.Cookies(request.URL)
	
	return data, response.StatusCode, err
}

/**
 * CURL POST
 */
func (self *CurlBase) Post(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	sensdData := self.encodeParams(params)
	apiUrl := requestUrl

	request, err := http.NewRequest("POST", apiUrl, strings.NewReader(sensdData))
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"NewRequest Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Api-Token", self.ApiToken)
	//request.Header.Add("Connection", "close")
	//self.setRequestCookie(request)
	response, err := self.client.Do(request)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"client.Do Failed",
			apiUrl,
			err.Error(),
		)
		
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	defer response.Body.Close()

	//保存 cookie
	/*respCks := response.Cookies()
	fmt.Println("respCks", respCks)
	appendCookie := self.cookies
	self.cookies = append(appendCookie, respCks...)*/
	
	data, err := ioutil.ReadAll(response.Body)
	HelperLog.TraceLog("POST SEND: " + apiUrl + sensdData)
	HelperLog.TraceLog("CURL POST : " + string(data))
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	//gCurCookies = gCurCookieJar.Cookies(request.URL)
	
	return data, response.StatusCode, err
}


/**
 * CURL MultiPart Put
 */
func (self *CurlBase) MultiPartPut(requestUrl string, params *bytes.Buffer, sendContentType string) ([]byte, int, error) {
	
	apiUrl := requestUrl
	request, err := http.NewRequest("PUT", apiUrl, params)
	HelperLog.TraceLog("MultiPartPut SEND: " + apiUrl + params.String())
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"NewRequest Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	request.Header.Set("Content-Type", sendContentType)
	request.Header.Add("Api-Token", self.ApiToken)
	//request.Header.Add("Connection", "close")
	//self.setRequestCookie(request)
	response, err := self.client.Do(request)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"client.Do Failed",
			apiUrl,
			err.Error(),
		)
		
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	defer response.Body.Close()
	
	//保存 cookie
	/*respCks := response.Cookies()
	fmt.Println("respCks", respCks)
	appendCookie := self.cookies
	self.cookies = append(appendCookie, respCks...)*/
	
	data, err := ioutil.ReadAll(response.Body)
	
	HelperLog.TraceLog("CURL MultiPartPost : " + string(data))
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	//gCurCookies = gCurCookieJar.Cookies(request.URL)
	
	return data, response.StatusCode, err
}

/**
 * CURL PUT
 */
func (self *CurlBase) Put(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	sensdData := self.encodeParams(params)
	apiUrl := requestUrl

	request, err := http.NewRequest("PUT", apiUrl, strings.NewReader(sensdData))
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"NewRequest Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Api-Token", self.ApiToken)
	request.Header.Add("Connection", "close")
	//self.setRequestCookie(request)
	response, err := self.client.Do(request)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"client.Do Failed",
			apiUrl,
			err.Error(),
		)
		
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	defer response.Body.Close()
	
	//保存 cookie
	/*respCks := response.Cookies()
	fmt.Println("respCks", respCks)
	appendCookie := self.cookies
	self.cookies = append(appendCookie, respCks...)*/
	
	data, err := ioutil.ReadAll(response.Body)
	HelperLog.TraceLog("PUT SEND: " + apiUrl + sensdData)
	HelperLog.TraceLog("CURL PUT : " + string(data))
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	//gCurCookies = gCurCookieJar.Cookies(request.URL)
	
	return data, response.StatusCode, err
}

/**
 * CURL DELETE
 */
func (self *CurlBase) Delete(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	sensdData := self.encodeParams(params)
	apiUrl := requestUrl

	request, err := http.NewRequest("DELETE", apiUrl, strings.NewReader(sensdData))
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"NewRequest Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Api-Token", self.ApiToken)
	request.Header.Add("Connection", "close")
	//self.setRequestCookie(request)
	response, err := self.client.Do(request)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"client.Do Failed",
			apiUrl,
			err.Error(),
		)
		
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	defer response.Body.Close()

	//保存 cookie
	/*respCks := response.Cookies()
	appendCookie := self.cookies
	self.cookies = append(appendCookie, respCks...)*/
	
	data, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	//gCurCookies = gCurCookieJar.Cookies(request.URL)
	
	return data, response.StatusCode, err
}

/**
 * 設定Request Cookie
 */
func (self *CurlBase) setRequestCookie(request *http.Request)  {
	fmt.Println(request)
	fmt.Println("self.cookies", self.cookies)
	for _, v := range self.cookies {
		request.AddCookie(v)
	}
}

/**
 * 所有參數轉為uri格式字串,參數依照key值排序
 * ex. hash=asd&param1=a&param2=b
 */
func (self *CurlBase) encodeParams(params map[string]interface{}) string {
	u := url.URL{}
	
	//add paramter
	query := u.Query()
	for k, v := range params {
		query.Set(k, fmt.Sprint(v))
	}
	
	return query.Encode()
}