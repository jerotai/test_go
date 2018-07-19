package apicurl

import (
	"Stingray/trait"
	"github.com/gin-gonic/gin"
	"Stingray/helper"
)

/**
 * Site Post Api
 */
func (a *apiNormalCurl) SiteSendPost(g *gin.Context) {
	a.sendPost(g, trait.Site_Code)
}

/**
 * Hall Post Api
 */
func (a *apiNormalCurl) HallSendPost(g *gin.Context) {
	a.sendPost(g, trait.Hall_Code)
}

/**
 * Site Put Api
 */
func (a *apiNormalCurl) SiteSendPut(g *gin.Context) {
	a.sendPut(g, trait.Site_Code)
}

/**
 * Hall Put Api
 */
func (a *apiNormalCurl) HallSendPut(g *gin.Context) {
	a.sendPut(g, trait.Hall_Code)
}

/**
 * Site Delete Api
 */
func (a *apiNormalCurl) SiteSendDelete(g *gin.Context) {
	a.sendDelete(g, trait.Site_Code)
}

/**
 * Hall Delete Api
 */
func (a *apiNormalCurl) HallSendDelete(g *gin.Context) {
	a.sendDelete(g, trait.Hall_Code)
}

/**
 * Send Post Api
 */
func (a *apiNormalCurl) sendPost(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[sendPost DataParse Error] " + apiRequestUrl + ":" + err.Error())
	}
	
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
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Send Put Api
 */
func (a *apiNormalCurl) sendPut(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPutApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[sendPut DataParse Error] " + apiRequestUrl + ":" + err.Error())
	}
	
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
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Send Delete Api
 */
func (a *apiNormalCurl) sendDelete(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitDeleteApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[sendDelete DataParse Error] " + apiRequestUrl + ":" + err.Error())
	}
	
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
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}
