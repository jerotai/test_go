package jellyfish

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"Stingray/api/apiparse"
	"strings"
)

/**
 * JellyFish Get Api
 */
func (j *JellyFish) HallSendGet(g *gin.Context) {
	//check api url path
	urlSplit := strings.Split(g.Request.URL.Path, "/")
	//check api conf
	j.InitGetApiConfig("/" + urlSplit[1] + "/" + urlSplit[2])
	
	inputDto := j.ApiConf.dto
	
	apiparse.GetDataParse(g, inputDto)
	sendParams, _ := apiparse.TowLayerDtoToMap(inputDto)
	
	//get hall code
	sendParams["hall_code"] = apiparse.StationCodeParse(g.Request)
	
	//call jellyFish service api
	ApiToken := j.GetToken(g)
	j.SendCurl.ApiToken = ApiToken
	body, httpStatus, err := j.SendCurl.Get(j.ApiConf.apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}
	
	g.String(httpStatus, string(body))
}

/**
 * JellyFish Post Api
 */
func (j *JellyFish) HallSendPost(g *gin.Context) {
	//check api conf
	j.InitPostApiConfig(g.Request.URL.Path)
	
	inputDto := j.ApiConf.dto
	
	apiparse.PostDataParse(g, inputDto)
	fmt.Println("inputDto: ", inputDto)
	sendParams, _ := apiparse.TowLayerDtoToMap(inputDto)
	
	ApiToken := j.GetToken(g)
	j.SendCurl.ApiToken = ApiToken
	
	//get hall code
	sendParams["hall_code"] = apiparse.StationCodeParse(g.Request)
	
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Post(j.ApiConf.apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}
	
	fmt.Println(string(body))
	g.String(httpStatus, string(body))
}

/**
 * JellyFish Put Api
 */
func (j *JellyFish) HallSendPut(g *gin.Context) {
	//check api conf
	j.InitPutApiConfig(g.Request.URL.Path)
	
	inputDto := j.ApiConf.dto
	
	apiparse.RsaDataParse(g, inputDto)
	fmt.Println("inputDto: ", inputDto)
	sendParams, _ := apiparse.TowLayerDtoToMap(inputDto)
	
	ApiToken := j.GetToken(g)
	j.SendCurl.ApiToken = ApiToken
	
	//get hall code
	sendParams["hall_code"] = apiparse.StationCodeParse(g.Request)
	
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Put(j.ApiConf.apiRequestUrl, sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}
	
	fmt.Println(string(body))
	g.String(httpStatus, string(body))
}


