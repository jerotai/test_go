package apicurl

import (
	dto "Stingray/core/dto/whitebaitdto"
	
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"encoding/base64"
	"encoding/json"
	"Stingray/trait"
)

/**
 * USer LoginAccount
 */
func (a *apiCurl) WhitebaitLogin(g *gin.Context) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[WhitebaitLogin Error] " + apiRequestUrl + ":" + err.Error())
	}
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get site code
	sendParams[trait.Site_Code] = trait.StationCodeParse(g.Request)
	sendParams["login_ip"] = trait.GetUserIp(g)
	sendParams["domain"] = trait.GetUserDomain(g.Request, true)
	
	//call whitebait service api
	body, httpStatus, err := a.SendCurl.Post(a.apiServiceUrl + "user/login", sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[Whitebait Login] " + err.Error())
	}
	
	// return rsa public key to client
	jsonMap := make(map[string]interface{})
	loginRes := dto.UserLoginRes{}
	json.Unmarshal(body, &loginRes)
	
	result := make(map[string]interface{})
	result["api_token"] =  loginRes.Result.Api_Token
	result["key"] = ""
	
	if loginRes.Code == "1" {
		rsaPublicKey, rsaPrivateKey := helper.CreateRsaKey()
		trait.SaveLoginRsaKey(loginRes.Result.Api_Token, rsaPublicKey, rsaPrivateKey)
		
		base64Code := base64.StdEncoding.EncodeToString([]byte(rsaPublicKey))
		result["key"] = base64Code
	}
	
	jsonMap["code"] = loginRes.Code
	jsonMap["result"] = result
	
	if loginRes.Message != "" {
		jsonMap["message"] = loginRes.Message
	}

	g.JSON(httpStatus, jsonMap)
}

/**
 * User Register
 */
func (a *apiCurl) WhitebaitRegister (g *gin.Context) {
	//check api conf
	apiRequestUrl, inputDto := a.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[WhitebaitRegister Error] " + apiRequestUrl + ":" + err.Error())
	}
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get site code
	sendParams[trait.Site_Code] = trait.StationCodeParse(g.Request)
	
	//call service api
	body, httpStatus, err := a.SendCurl.Post(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[Whitebait Register Api] " + err.Error())
	}
	
	trait.UpdateRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}