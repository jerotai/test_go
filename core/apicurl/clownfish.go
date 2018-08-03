package apicurl

import (
	"github.com/gin-gonic/gin"
	"Stingray/trait"
	"Stingray/helper"
	"strings"
)

/**
 * USer LoginAccount
 */
func (a *apiServiceCurl) CooperationFortune(g *gin.Context) {
	//check api url path
	urlSplit := strings.Split(g.Request.URL.Path, "/")
	
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitGetApiConfig("/" + urlSplit[1] + "/" + urlSplit[2])
	
	sendParams, _ := trait.DataParseByGet(g, inputDto)
	
	//get station type code
	sendParams[trait.Site_Code] = trait.StationCodeParse(g.Request)
	
	//call service api
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	body, httpStatus, err := a.SendCurl.Get(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[CooperationFortune Api] " + apiRequestUrl + ":" + err.Error())
	}
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}
