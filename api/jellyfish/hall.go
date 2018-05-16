package jellyfish

import (
	"github.com/gin-gonic/gin"
	"Stingray/core/dto"
	"Stingray/api/apiparse"
	"fmt"
)

/**
 * 建立下層廳主
 */
func (j *JellyFish) Hall(g *gin.Context) {
	inputDto := dto.Hall{}
	apiparse.PostRsaDataParse(g, &inputDto)
	
	//get hall code
	inputDto.HallCode = apiparse.SiteCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMapInterface(inputDto)
	j.SendCurl.ApiToken = g.Request.Header.Get("Api-Token")
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Post("hall", sendParams)
	if err != nil {
		//todo 錯誤處理
	}
	
	g.String(httpStatus, string(body))
	
}

/**
 * 更新下層廳主資料
 */
func (j *JellyFish) UpdataHall(g *gin.Context) {
	inputDto := dto.UpdataHall{}
	apiparse.PostRsaDataParse(g, &inputDto)
	
	//get hall code
	inputDto.HallCode = apiparse.SiteCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMapInterface(inputDto)
	fmt.Println(sendParams)
	//call jellyFish service api
	j.SendCurl.ApiToken = g.Request.Header.Get("Api-Token")
	body, httpStatus, err := j.SendCurl.Put("hall", sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}
	
	g.String(httpStatus, string(body))
}

/**
 * 啟用下層廳主
 */
func (j *JellyFish) EnabledHall(g *gin.Context) {
	inputDto := dto.EnabledHall{}
	apiparse.PostRsaDataParse(g, &inputDto)
	
	//get hall code
	inputDto.HallCode = apiparse.SiteCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMapInterface(inputDto)
	
	//call jellyFish service api
	j.SendCurl.ApiToken = g.Request.Header.Get("Api-Token")
	body, httpStatus, err := j.SendCurl.Put("hall/enabled", sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}

	g.String(httpStatus, string(body))
}

func (j *JellyFish) SubList(g *gin.Context) {
	inputDto := dto.SubList{}
	apiparse.GetDataParse(g, &inputDto)
	//get hall code
	inputDto.HallCode = apiparse.SiteCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMapInterface(inputDto)
	//call jellyFish service api
	j.SendCurl.ApiToken = g.Request.Header.Get("Api-Token")
	body, httpStatus, err := j.SendCurl.Get("hall/subList", sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}
	g.String(httpStatus, string(body))
}