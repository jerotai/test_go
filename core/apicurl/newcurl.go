package apicurl

import (
	"Stingray/trait"
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"strings"
)

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
 * Send Get Api
 */
func (a *apiCurl) sendGet(g *gin.Context, stationType string) {
	//check api url path
	urlSplit := strings.Split(g.Request.URL.Path, "/")
	
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitGetApiConfig("/" + urlSplit[1] + "/" + urlSplit[2])
	
	sendParams, _ := trait.DataParseByGet(g, inputDto)
	//sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
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
	
	trait.UpdateRsaKeyExpire(ApiToken)
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
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}
