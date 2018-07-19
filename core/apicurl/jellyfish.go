package apicurl

import (
	"github.com/gin-gonic/gin"
	dto "Stingray/core/dto/jellyfishdto"
	"Stingray/helper"
	"encoding/json"
	"encoding/base64"
	"Stingray/trait"
)

/**
 * JellyFish Login Api
 */
func (j *apiCurl) JellyFishLogin(g *gin.Context) {
	//check api conf
	apiRequestUrl, inputDto := j.apiConfInit.InitPostApiConfig(g.Request.URL.Path)
	
	sendParams, err := trait.DataParse(g, inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[JellyFish-Login DataParse Error] " + apiRequestUrl + ":" + err.Error())
	}
	
	sendParams[trait.Hall_Code] = trait.StationCodeParse(g.Request)
	
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Post(j.apiServiceUrl + apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[JellyFish-Login] " + err.Error())
		panic(err)
	}
	
	// return rsa public key to client
	jsonMap := make(map[string]interface{})
	loginRes := dto.LoginRes{}
	json.Unmarshal(body, &loginRes)
	result := make(map[string]interface{})
	result["api_token"] =  loginRes.Result.Api_Token
	result["role_id"] =  loginRes.Result.Role_Id
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
