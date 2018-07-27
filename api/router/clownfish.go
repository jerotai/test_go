package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/api/clownfish"
	"Stingray/helper"
	"Stingray/core/apicurl"
)

func InitClownfishRouting(routerGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("clownfish_service")
	router := clownfish.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	var apiCurlSend, apiServiceSend = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * cooperation Group
	 */
	cooperation := routerGroup.Group("/cooperation")
	{
		cooperation.GET("/fortune/:Code", apiServiceSend.CooperationFortune)
		cooperation.GET("/fortuneUser/:Code", apiCurlSend.SiteSendGet)
	}
}
