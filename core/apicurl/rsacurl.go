package apicurl

import (
	"Stingray/trait"
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"net/http"
)

/**
 * Site Post Api
 */
func (a *apiRsaCurl) SiteSendPost(g *gin.Context) {
	a.sendPost(g, trait.Site_Code)
}

/**
 * Hall Post Api
 */
func (a *apiRsaCurl) HallSendPost(g *gin.Context) {
	a.sendPost(g, trait.Hall_Code)
}

/**
 * Site Put Api
 */
func (a *apiRsaCurl) SiteSendPut(g *gin.Context) {
	a.sendPut(g, trait.Site_Code)
}

/**
 * Hall Put Api
 */
func (a *apiRsaCurl) HallSendPut(g *gin.Context) {
	a.sendPut(g, trait.Hall_Code)
}

/**
 * Site Delete Api
 */
func (a *apiRsaCurl) SiteSendDelete(g *gin.Context) {
	a.sendDelete(g, trait.Site_Code)
}

/**
 * Hall Delete Api
 */
func (a *apiRsaCurl) HallSendDelete(g *gin.Context) {
	a.sendDelete(g, trait.Hall_Code)
}

/**
 * Send Post Api
 */
func (a *apiRsaCurl) sendPost(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParseByRsa(g, inputDto)
	if err != nil {
		reBody := trait.ApiResponse{}
		reBody.Code = trait.RSA_KEY_EMPTY
		reBody.Message = err.Error()
		g.JSON(http.StatusOK, reBody)
	} else {
		//sendParams, _ := trait.TowLayerDtoToMap(inputDto)
		
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
}

/**
 * Send Put Api
 */
func (a *apiRsaCurl) sendPut(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPutApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParseByRsa(g, inputDto)
	if err != nil {
		reBody := trait.ApiResponse{}
		reBody.Code = trait.RSA_KEY_EMPTY
		reBody.Message = err.Error()
		g.JSON(http.StatusOK, reBody)
	} else {
		//sendParams, _ := trait.TowLayerDtoToMap(inputDto)
		
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
}

/**
 * Send Delete Api
 */
func (a *apiRsaCurl) sendDelete(g *gin.Context, stationType string) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitDeleteApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParseByRsa(g, inputDto)
	if err != nil {
		reBody := trait.ApiResponse{}
		reBody.Code = trait.RSA_KEY_EMPTY
		reBody.Message = err.Error()
		g.JSON(http.StatusOK, reBody)
	} else {
		//sendParams, _ := trait.TowLayerDtoToMap(inputDto)
		
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
}
