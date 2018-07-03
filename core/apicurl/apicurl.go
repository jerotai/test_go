package apicurl

import (
	"github.com/gin-gonic/gin"
	"strings"
	"Stingray/helper"
	"Stingray/trait"
	"fmt"
	"sync"
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

/**
 * Site Get Api
 */
func (a *apiCurl) SiteSendGet(g *gin.Context) {
	a.sendGet(g, trait.Site_Code)
}

/**
 * Hall Get Api
 */
func (a *apiCurl) HallSendGet(g *gin.Context) {
	a.sendGet(g, trait.Hall_Code)
}

/**
 * Site Post Api
 */
func (a *apiCurl) SiteSendPost(g *gin.Context) {
	a.sendPost(g, trait.Site_Code)
}

/**
 * Hall Post Api
 */
func (a *apiCurl) HallSendPost(g *gin.Context) {
	a.sendPost(g, trait.Hall_Code)
}

/**
 * Site MultiPartPost Api
 */
func (a *apiCurl) SiteSendMultiPartPost(g *gin.Context) {
	a.sendMultiPartPost(g, trait.Site_Code)
}

/**
 * Hall MultiPartPost Api
 */
func (a *apiCurl) HallSendMultiPartPost(g *gin.Context) {
	a.sendMultiPartPost(g, trait.Hall_Code)
}

/**
 * Site Put Api
 */
func (a *apiCurl) SiteSendPut(g *gin.Context) {
	a.sendPut(g, trait.Site_Code)
}

/**
 * Hall Put Api
 */
func (a *apiCurl) HallSendPut(g *gin.Context) {
	a.sendPut(g, trait.Hall_Code)
}

/**
 * Site Delete Api
 */
func (a *apiCurl) SiteSendDelete(g *gin.Context) {
	a.sendDelete(g, trait.Site_Code)
}

/**
 * Hall Delete Api
 */
func (a *apiCurl) HallSendDelete(g *gin.Context) {
	a.sendDelete(g, trait.Hall_Code)
}

/**
 * Send Get Api
 */
func (a *apiCurl) sendGet(g *gin.Context, stationType string) {
	//check api url path
	urlSplit := strings.Split(g.Request.URL.Path, "/")
	
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitGetApiConfig("/" + urlSplit[1] + "/" + urlSplit[2])
	
	trait.DataParseByGet(g, inputDto)
	sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
	//get station type code
	sendParams[stationType] = trait.StationCodeParse(g.Request)
	
	//call service api
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	body, httpStatus, err := a.SendCurl.Get(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[SendGet Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdataRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Send Post Api
 */
func (a *apiCurl) sendPost(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	trait.DataParse(g, inputDto)
	/*
	err := trait.DataParseByRsa(g, inputDto)
	if err != nil {
		reBody := trait.ApiResponse{}
		reBody.Code = trait.RSA_KEY_EMPTY
		reBody.Message = err.Error()
		g.JSON(http.StatusOK, reBody)
	}
	*/

	sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get station type code
	sendParams[stationType] = trait.StationCodeParse(g.Request)

	//call service api
	body, httpStatus, err := a.SendCurl.Post(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[SendPost Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdataRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Send MultiPartPost Api
 */
func (a *apiCurl) sendMultiPartPost(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	sendParams, sendContentType := trait.MultiPartDataParse(g, inputDto, stationType, trait.StationCodeParse(g.Request))
	
	//call service api
	body, httpStatus, err := a.SendCurl.MultiPartPost(a.apiServiceUrl + apiRequestUrl, sendParams, sendContentType)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[SendMultiPartPost Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdataRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Send Put Api
 */
func (a *apiCurl) sendPut(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPutApiConfig(g.Request.URL.Path)
	
	trait.DataParse(g, inputDto)
	/*
	err := trait.DataParseByRsa(g, inputDto)
	if err != nil {
		reBody := trait.ApiResponse{}
		reBody.Code = trait.RSA_KEY_EMPTY
		reBody.Message = err.Error()
		g.JSON(http.StatusOK, reBody)
	}
	*/
	
	sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get station type code
	sendParams[stationType] = trait.StationCodeParse(g.Request)
	
	//call service api
	body, httpStatus, err := a.SendCurl.Put(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[SendPut Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdataRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Send Delete Api
 */
func (a *apiCurl) sendDelete(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitDeleteApiConfig(g.Request.URL.Path)
	
	trait.DataParse(g, inputDto)
	/*
	err := trait.DataParseByRsa(g, inputDto)
	if err != nil {
		reBody := trait.ApiResponse{}
		reBody.Code = trait.RSA_KEY_EMPTY
		reBody.Message = err.Error()
		g.JSON(http.StatusOK, reBody)
	}
	*/
	
	sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get station type code
	sendParams[stationType] = trait.StationCodeParse(g.Request)
	
	//call service api
	body, httpStatus, err := a.SendCurl.Delete(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[SendDelete Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdataRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}


