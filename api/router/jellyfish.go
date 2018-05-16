package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/api/jellyfish"
	"Stingray/helper"
)

func InitJellyFishRouting(jellyfishGroup *gin.RouterGroup) {
	jellyfish := jellyfish.New()
	jellyfish.SendCurl = helper.NewCurlBase()
	
	apiConf := helper.ApiSetting("JellyFishService")
	jellyfish.SendCurl.Url = apiConf.Host + ":" + apiConf.Port + "/"
	
	/**
	 * User Group
	 */
	user := jellyfishGroup.Group("/login")
	{
		user.POST((""), jellyfish.Login)
	}
	
	/**
	 * Hall Group
	 */
	hall := jellyfishGroup.Group("/hall")
	{
		hall.POST((""), jellyfish.Hall)
		hall.PUT((""), jellyfish.UpdataHall)
		hall.PUT(("/enabled"), jellyfish.EnabledHall)
		hall.GET(("/subList/:Page/:Count"), jellyfish.SubList)
	}
	
}
