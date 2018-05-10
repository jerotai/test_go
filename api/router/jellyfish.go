package router

import (
	"routes/api/jellyfish"
	"github.com/gin-gonic/gin"
	"routes/helper"
)


func InitJellyFishRouting(jellyfishGroup *gin.RouterGroup) {
	jellyfish := jellyfish.New()
	jellyfish.SendCurl = &helper.CurlBase{}
	apiConf := helper.ApiSetting("JellyFishService")
	jellyfish.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	
	jellyfishGroup.POST(("/login"), jellyfish.Login)
}
