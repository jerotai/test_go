package jellyfish


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"routes/core/dto"
	"routes/api/apiparse"
)

/**
 * JellyFish Login Api
 */
func (j *JellyFish) Login(g *gin.Context) {
	inputDto := dto.Login{}
	apiparse.PostRsaDataParse(g, &inputDto)
	
	//get hall code
	inputDto.HallCode = apiparse.SiteCodeParse(g.Request)
	sendParams, _ := apiparse.DtoToMapInterface(inputDto)
	
	response := &dto.LoginResponse{}
	
	//call jellyFish service api
	err := j.SendCurl.Post("login", sendParams, &response)
	
	if err != nil {
		//todo 錯誤處理
	}
	
	g.JSON(http.StatusOK, inputDto)
}

