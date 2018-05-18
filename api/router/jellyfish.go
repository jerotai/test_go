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
	 * Hall Group (廳主)
	 */
	hall := jellyfishGroup.Group("/hall")
	{
		hall.POST((""), jellyfish.HallSendPost)
		hall.PUT((""), jellyfish.HallSendPut)
		hall.PUT(("/enabled"), jellyfish.HallSendPut)
		hall.GET(("/subList/:Page/:Count"), jellyfish.HallSendGet)
	}
	
	/**
	 * Shareholder Group (股東)
	 */
	shareholder := jellyfishGroup.Group("/shareholder")
	{
		//hall_code=CQ1&site_code=AA01&status=1&name=AA01
		//&start_time=2018-05-01 00:00:00&end_time=2018-05-12 23:59:59
		shareholder.GET(("/list/:Site_Code/:Page/:Count/"), jellyfish.HallSendGet)
		shareholder.GET(("/data/:Id"), jellyfish.HallSendGet)
		shareholder.POST((""), jellyfish.HallSendPost)
		shareholder.PUT((""), jellyfish.HallSendPut)
	}
	
	/**
	 * hallSub Grpup (廳主子帳號)
	 */
	hallSub := jellyfishGroup.Group("/hallSub")
	{
		hallSub.POST((""), jellyfish.HallSendPost)
		hallSub.PUT("", jellyfish.HallSendPut)
		hallSub.GET("/list/:Page/:Count",jellyfish.HallSendGet)
		
	}
	
	/**
	 * authGroup Grpup (權限群組)
	 */
	authgroup := jellyfishGroup.Group("authgroup")
	{
		authgroup.POST("", jellyfish.HallSendPost)
		authgroup.PUT((""), jellyfish.HallSendPut)
		authgroup.GET("/list/:Page/:Count",jellyfish.HallSendGet)
		authgroup.GET("/detail/:Group_Id",jellyfish.HallSendGet)
		authgroup.GET("/dropDown",jellyfish.HallSendGet)
	}
}
