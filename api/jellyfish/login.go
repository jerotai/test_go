package jellyfish

import (
	"github.com/gin-gonic/gin"
	"Stingray/core/dto"
	"Stingray/api/apiparse"
)

/**
 * JellyFish Login Api
 */
func (j *JellyFish) Login(g *gin.Context) {
	inputDto := dto.Login{}
	apiparse.PostDataParse(g, &inputDto)
	
	//get hall code
	inputDto.HallCode = apiparse.SiteCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMapInterface(inputDto)
	
	//call jellyFish service api
	body, httpStatus, err := j.SendCurl.Post("login", sendParams)
	
	if err != nil {
		//todo 錯誤處理
	}
	
	g.String(httpStatus, string(body))
}

