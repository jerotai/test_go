package apicurl

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/trait"
)

/**
 * Create Article Api
 */
func (a *apiServiceCurl) CreateArticle(g *gin.Context) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[createArticle DataParse Error] " + apiRequestUrl + ":" + err.Error())
	}
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get station type code
	sendParams[trait.Hall_Code] = trait.StationCodeParse(g.Request)
	
	//call service api
	body, httpStatus, err := a.SendCurl.Post(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[createArticle Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}

/**
 * Update Article Api
 */
func (a *apiServiceCurl) UpdateArticle(g *gin.Context) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPutApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[UpdateArticle DataParse Error] " + apiRequestUrl + ":" + err.Error())
	}
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get station type code
	sendParams[trait.Hall_Code] = trait.StationCodeParse(g.Request)
	
	//call service api
	body, httpStatus, err := a.SendCurl.Put(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[UpdateArticle Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}