package helper

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
	"fmt"
	"errors"
)

type CurlBase struct {
	cookies []*http.Cookie
	client *http.Client
	ApiToken string
	Url  string
}

//初始化
func NewCurlBase() *CurlBase {
	hc := &CurlBase{}
	hc.client = &http.Client{}
	//為所有重定向的請求增加cookie
	hc.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			for _,v := range hc.GetCookie() {
				req.AddCookie(v)
			}
		}
		return nil
	}
	return hc
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
	self.cookies = append(self.cookies, cookies...)
}

//獲取當前所有的cookie
func (self *CurlBase) GetCookie() ([]*http.Cookie) {
	return self.cookies
}

/**
 * CURL GET
 */
func (self *CurlBase) Get(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	paramUri := self.encodeParams(params)

	apiUrl := self.Url + requestUrl + "?" + paramUri
	fmt.Println("send Get :", apiUrl)
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
	self.setRequestCookie(request)
	response,_ := self.client.Do(request)
	defer response.Body.Close()
	
	data, err := ioutil.ReadAll(response.Body)
	fmt.Println(data)
	if err != nil {
		errStr := fmt.Sprintf(
			"Msg:%s api:%s err:%s",
			"Read From body Failed",
			apiUrl,
			err.Error(),
		)
		return nil, http.StatusOK, errors.New(errStr)
	}
	
	return data, response.StatusCode, err
}

/**
 * CURL POST
 */
func (self *CurlBase) Post(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	postData := self.encodeParams(params)
	apiUrl := self.Url + requestUrl
	fmt.Println("send Post :", postData)
	request, err := http.NewRequest("POST", apiUrl, strings.NewReader(postData))
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
	self.setRequestCookie(request)
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
	fmt.Println(response.Status)
	defer response.Body.Close()
	
	//保存 cookie
	respCks := response.Cookies()
	self.cookies = append(self.cookies, respCks...)
	
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
	
	return data, response.StatusCode, err
}

/**
 * CURL PUT
 */
func (self *CurlBase) Put(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	postData := self.encodeParams(params)
	apiUrl := self.Url + requestUrl
	request, err := http.NewRequest("PUT", apiUrl, strings.NewReader(postData))
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
	self.setRequestCookie(request)
	
	response, err := self.client.Do(request)
	defer response.Body.Close()
	
	//保存响应的 cookie
	respCks := response.Cookies()
	self.cookies = append(self.cookies, respCks...)
	
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
	
	return data, response.StatusCode, err
}

/**
 * CURL DELETE
 */
func (self *CurlBase) Delete(requestUrl string, params map[string]interface{}) ([]byte, int, error) {
	postData := self.encodeParams(params)
	apiUrl := self.Url + requestUrl
	request, err := http.NewRequest("DELETE", apiUrl, strings.NewReader(postData))
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
	self.setRequestCookie(request)
	
	response, err := self.client.Do(request)
	defer response.Body.Close()
	
	//保存响应的 cookie
	respCks := response.Cookies()
	self.cookies = append(self.cookies, respCks...)
	
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
	
	return data, response.StatusCode, err
}

/**
 * 設定Request Cookie
 */
func (self *CurlBase) setRequestCookie(request *http.Request)  {
	for _,v := range self.cookies{
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