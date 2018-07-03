package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/api/stingray"
	"Stingray/core/apicurl"
)

func InitStingrayRouting(group *gin.RouterGroup) {
	router := stingray.New()
	apiConf := helper.ApiSetting("whitebait_service")
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)
	
	
	/**
	 * rsa/key
	 */
	group.POST("/rsa/key", apicurlSend.NewRsaPubKey)
}
