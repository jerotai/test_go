package apicurl

import (
	"Stingray/helper"
	"fmt"
	"sync"
	"github.com/gin-gonic/gin"
)

var (
	// 單例模式實例
	apiCurlInst *helper.CurlBase
	apiCurlOnce sync.Once
)

func apiBaseCurl() *helper.CurlBase {
	apiCurlOnce.Do(func(){
		apiCurlInst = helper.NewCurlBase()
	})
	
	return apiCurlInst
}

type (
	// apiCurl 建構子
	apiCurl struct {
		SendCurl *helper.CurlBase
		apiServiceUrl string
		apiConfInit *ApiConfInit
	}
	apiRsaCurl struct {
		SendCurl *helper.CurlBase
		apiServiceUrl string
		apiConfInit *ApiConfInit
	}
	apiNormalCurl struct {
		SendCurl *helper.CurlBase
		apiServiceUrl string
		apiConfInit *ApiConfInit
	}
	apiServiceCurl struct {
		SendCurl *helper.CurlBase
		apiServiceUrl string
		apiConfInit *ApiConfInit
	}
)

type ApiConfInit struct {
	InitGetApiConfig func(string) (string, interface{})
	InitPostApiConfig func(string) (string, interface{})
	InitPutApiConfig func(string) (string, interface{})
	InitDeleteApiConfig func(string) (string, interface{})
}

func New(apiConf *helper.ApiSettingConf, apiConfInit *ApiConfInit) *apiCurl {
	route := &apiCurl{}
	route.SendCurl = apiBaseCurl()
	route.apiServiceUrl = fmt.Sprintf("%s:%d/", apiConf.Host, apiConf.Port)
	route.apiConfInit = apiConfInit
	return route
}

func NewRsa(apiConf *helper.ApiSettingConf, apiConfInit *ApiConfInit) *apiRsaCurl {
	route := &apiRsaCurl{}
	route.SendCurl = apiBaseCurl()
	route.apiServiceUrl = fmt.Sprintf("%s:%d/", apiConf.Host, apiConf.Port)
	route.apiConfInit = apiConfInit
	return route
}

func NewNormal(apiConf *helper.ApiSettingConf, apiConfInit *ApiConfInit) *apiNormalCurl {
	route := &apiNormalCurl{}
	route.SendCurl = apiBaseCurl()
	route.apiServiceUrl = fmt.Sprintf("%s:%d/", apiConf.Host, apiConf.Port)
	route.apiConfInit = apiConfInit
	return route
}

func NewService(apiConf *helper.ApiSettingConf, apiConfInit *ApiConfInit) *apiServiceCurl {
	route := &apiServiceCurl{}
	route.SendCurl = apiBaseCurl()
	route.apiServiceUrl = fmt.Sprintf("%s:%d/", apiConf.Host, apiConf.Port)
	route.apiConfInit = apiConfInit
	return route
}

type ApiCurlSendInit struct {
	SiteSendGet func(*gin.Context)
	HallSendGet func(*gin.Context)
	SiteSendPost func(*gin.Context)
	HallSendPost func(*gin.Context)
	SiteSendMultiPartPost func(*gin.Context)
	HallSendMultiPartPost func(*gin.Context)
	SiteSendPut func(*gin.Context)
	HallSendPut func(*gin.Context)
	SiteSendDelete func(*gin.Context)
	HallSendDelete func(*gin.Context)
}

/**
 * 回傳項目:
 * ApiCurlSendInit : 一般用CRUL 和 RSA用CURL (GET, POST, PUT, DELETE, MultiPartPost)
 * apiCurl : 各服務專用連線項目
 */
func GetCurlSend(apiConf *helper.ApiSettingConf, apiConfInit *ApiConfInit) (*ApiCurlSendInit, *apiServiceCurl) {
	apiServiceSend := NewService(apiConf, apiConfInit)
	apiNewSend := New(apiConf, apiConfInit)
	var apiCurlSend = &ApiCurlSendInit{}
	
	apiCurlSend.SiteSendGet = apiNewSend.SiteSendGet
	apiCurlSend.HallSendGet = apiNewSend.HallSendGet
	apiCurlSend.SiteSendMultiPartPost = apiNewSend.SiteSendMultiPartPost
	apiCurlSend.HallSendMultiPartPost = apiNewSend.HallSendMultiPartPost
	
	if helper.RsaOpen() == "Y" {
		curlSend := NewRsa(apiConf, apiConfInit)
		apiCurlSend.SiteSendPost = curlSend.SiteSendPost
		apiCurlSend.HallSendPost = curlSend.HallSendPost
		apiCurlSend.SiteSendPut = curlSend.SiteSendPut
		apiCurlSend.HallSendPut = curlSend.HallSendPut
		apiCurlSend.SiteSendDelete = curlSend.SiteSendDelete
		apiCurlSend.HallSendDelete = curlSend.HallSendDelete
	} else {
		curlSend := NewNormal(apiConf, apiConfInit)
		apiCurlSend.SiteSendPost = curlSend.SiteSendPost
		apiCurlSend.HallSendPost = curlSend.HallSendPost
		apiCurlSend.SiteSendPut = curlSend.SiteSendPut
		apiCurlSend.HallSendPut = curlSend.HallSendPut
		apiCurlSend.SiteSendDelete = curlSend.SiteSendDelete
		apiCurlSend.HallSendDelete = curlSend.HallSendDelete
	}
	
	return apiCurlSend, apiServiceSend
}