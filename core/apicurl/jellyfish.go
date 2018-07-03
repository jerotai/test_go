package apicurl

import (
	"github.com/gin-gonic/gin"
	dto "Stingray/core/dto/jellyfishdto"
	"Stingray/helper"
	"encoding/json"
	"encoding/base64"
	"Stingray/trait"
)


type LoginRes struct {
	Code string `json:"code"`
	Result struct{
		Api_Token string `json:"api_token"`
		Role_Id int `json:"role_id"`
	} `json:"result"`
	Message string `json:"message"`
}

/**
 * JellyFish Login Api
 */
func (j *apiCurl) JellyFishLogin(g *gin.Context) {
	inputDto := dto.Login{}
	trait.DataParse(g, &inputDto)
	
	//get hall code
	inputDto.Hall_Code = trait.StationCodeParse(g.Request)
	sendParams, err := trait.DtoToMap(inputDto)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[JellyFish-Login] DtoToMap " + err.Error())
	}
	
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Post(j.apiServiceUrl + "login", sendParams)
	
	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[JellyFish-Login] " + err.Error())
		panic(err)
	}
	
	// return rsa public key to client
	jsonMap := make(map[string]interface{})
	loginRes := LoginRes{}
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
