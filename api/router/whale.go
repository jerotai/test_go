package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/api/whale"
	"Stingray/helper"
	"Stingray/core/apicurl"
)


func InitWhaleRouting(whaleGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("whale_service")
	router := whale.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)
	
	/**
	 * Menu Group
	 */
	menu := whaleGroup.Group("/menu/list")
	{
		menu.GET("", apicurlSend.HallSendGet)
	}
	
	/**
	 * Site Group
	 */
	site := whaleGroup.Group("/site")
	{
		site.GET("/list/:Page/:Count/", apicurlSend.HallSendGet)
		site.GET("/dropdownlist/", apicurlSend.HallSendGet)
		site.GET("/data/:Id", apicurlSend.HallSendGet)
		
		site.POST("", apicurlSend.HallSendMultiPartPost)
		site.POST("/update", apicurlSend.HallSendMultiPartPost)
		
		site.PUT("/enabled", apicurlSend.HallSendPut)
		site.PUT("/agent/default", apicurlSend.HallSendPut)
	}
	
	/**
	 * Banner Group
	 */
	banner := whaleGroup.Group("/banner")
	{
		banner.GET("/list", apicurlSend.SiteSendGet)
		banner.GET("/backEndlist/:Site_Code", apicurlSend.HallSendGet)
		banner.GET("/data/:Site_Code/:Id", apicurlSend.HallSendGet)
		
		banner.POST("", apicurlSend.HallSendMultiPartPost)
		banner.POST("/update", apicurlSend.HallSendMultiPartPost)
		
		banner.DELETE("", apicurlSend.HallSendDelete)
	}
	
	/**
	 * promotion Group
	 */
	promotion := whaleGroup.Group("/promotion")
	{
		promotion.GET("/list", apicurlSend.SiteSendGet)
		promotion.GET("/info/:Id", apicurlSend.SiteSendGet)
		promotion.GET("/backEndList/:Site_Code", apicurlSend.HallSendGet)
		promotion.GET("/data/:Id/:Site_Code", apicurlSend.HallSendGet)
		
		promotion.POST("", apicurlSend.HallSendMultiPartPost)
		promotion.POST("/update", apicurlSend.HallSendMultiPartPost)
		
		promotion.DELETE("", apicurlSend.HallSendDelete)
	}
}
