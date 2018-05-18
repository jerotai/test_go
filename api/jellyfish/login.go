package jellyfish

import (
	"github.com/gin-gonic/gin"
	dto "Stingray/core/dto/jellyfishdto"
	"Stingray/api/apiparse"
	"Stingray/helper"
	"Stingray/core/model/redis/rsakey"
	"encoding/json"
	"encoding/base64"
)

type LoginRes struct {
	Code string `json:"code"`
	Result string `json:"result"`
	Message string `json:"message"`
}

/**
 * JellyFish Login Api
 */
func (j *JellyFish) Login(g *gin.Context) {
	inputDto := dto.Login{}
	apiparse.PostDataParse(g, &inputDto)
	
	//get hall code
	inputDto.Hall_Code = apiparse.StationCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMap(inputDto)
	
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Post("login", sendParams)
	
	if err != nil {
		//todo 錯誤處理
		panic(err)
	}
	
	apiResponse := apiparse.ApiResponse(string(body))
	rsaPublicKey, rsaPrivateKey := getRsaPublicKey(apiResponse.Result, inputDto.Account)
	status, err := saveLoginRsaKey(apiResponse.Result, rsaPublicKey, rsaPrivateKey)

	if (err != nil || status != "OK") {
		//todo 錯誤處理
		panic(err)
	}
	
	// return rsa public key to client
	jsonMap := make(map[string]interface{})
	logRes := LoginRes{}
	json.Unmarshal(body, &logRes)
	
	base64Code := base64.StdEncoding.EncodeToString([]byte(rsaPublicKey))
	result := make(map[string]string)
	result["token"] =  logRes.Result
	result["key"] = base64Code
	
	jsonMap["code"] = logRes.Code
	jsonMap["result"] = result
	
	if logRes.Message != "" {
		jsonMap["message"] = logRes.Message
	}
	
	g.JSON(httpStatus, jsonMap)
}

/**
 * get rsa keys
 */
func getRsaPublicKey(token string, account string) (string, string) {
	return helper.CreateRsaKey(token, account)
}

/**
 * save login rsa keys to redis
 */
func saveLoginRsaKey(token string, rsaPublicKey string, rsaPrivateKey string) (string, error) {
	
	rsaKeyRedis := rsakey.RsaKeyService()
	rsaKeyRedis.Init()
	
	return rsaKeyRedis.SetTokenRsaKey(token, rsaPublicKey, rsaPrivateKey)
}