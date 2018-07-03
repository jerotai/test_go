package apicurl

import (
	dto "Stingray/core/dto/whitebaitdto"
	
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"encoding/base64"
	"encoding/json"
	"Stingray/trait"
)

type UserLoginRes struct {
	Code string `json:"code"`
	Result struct{
		Api_Token string `json:"api_token"`
	} `json:"result"`
	Message string `json:"message"`
}

/**
 * USer LoginAccount
 */
func (a *apiCurl) WhitebaitLogin(g *gin.Context) {
	//check api conf
	inputDto := &dto.UserLogin{}
	trait.DataParse(g, inputDto)
	sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get hall code
	sendParams["site_code"] = trait.StationCodeParse(g.Request)
	sendParams["login_ip"] = trait.GetUserIp(g)
	
	//call whitebait service api
	body, httpStatus, err := a.SendCurl.Post(a.apiServiceUrl + "user/login", sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[Whitebait Login] " + err.Error())
	}
	
	// return rsa public key to client
	jsonMap := make(map[string]interface{})
	loginRes := UserLoginRes{}
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
	
	trait.DataParse(g, inputDto)

	sendParams, _ := trait.TowLayerDtoToMap(inputDto)
	
	ApiToken := trait.GetToken(g)
	a.SendCurl.ApiToken = ApiToken
	
	//get site code
	sendParams["site_code"] = trait.StationCodeParse(g.Request)
	
	//call service api
	body, httpStatus, err := a.SendCurl.Post(a.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[Whitebait Register Api] " + err.Error())
	}
	
	trait.UpdataRsaKeyExpire(ApiToken)
	g.String(httpStatus, string(body))
}