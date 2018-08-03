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
	
	var apiCurlSend, apiServiceSend = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * Menu Group
	 */
	menu := whaleGroup.Group("/menu/list")
	{
		menu.GET("", apiCurlSend.HallSendGet)
	}
	
	/**
	 * Site Group
	 */
	site := whaleGroup.Group("/site")
	{
		site.GET("/list/:Page/:Count/", apiCurlSend.HallSendGet)
		site.GET("/dropdownlist/", apiCurlSend.HallSendGet)
		site.GET("/data/:Id", apiCurlSend.HallSendGet)
		
		site.POST("", apiCurlSend.HallSendMultiPartPost)
		site.POST("/update", apiCurlSend.HallSendMultiPartPost)
		
		site.PUT("/enabled", apiCurlSend.HallSendPut)
		site.PUT("/agent/default", apiCurlSend.HallSendPut)
	}
	
	/**
	 * Banner Group
	 */
	banner := whaleGroup.Group("/banner")
	{
		banner.GET("/list", apiCurlSend.SiteSendGet)
		banner.GET("/backEndlist/:Site_Code", apiCurlSend.HallSendGet)
		banner.GET("/data/:Site_Code/:Id", apiCurlSend.HallSendGet)
		
		banner.POST("", apiCurlSend.HallSendMultiPartPost)
		banner.POST("/update", apiCurlSend.HallSendMultiPartPost)
		
		banner.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * promotion Group
	 */
	promotion := whaleGroup.Group("/promotion")
	{
		promotion.GET("/list", apiCurlSend.SiteSendGet)
		promotion.GET("/info/:Id", apiCurlSend.SiteSendGet)
		promotion.GET("/backEndList/:Site_Code", apiCurlSend.HallSendGet)
		promotion.GET("/data/:Id/:Site_Code", apiCurlSend.HallSendGet)
		
		promotion.POST("", apiCurlSend.HallSendMultiPartPost)
		promotion.POST("/update", apiCurlSend.HallSendMultiPartPost)
		
		promotion.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * article Group
	 */
	article := whaleGroup.Group("/article")
	{
		article.GET("/info/:Id", apiCurlSend.SiteSendGet)
		article.GET("/list/:Site_Code", apiCurlSend.HallSendGet)
		article.GET("/page/:Site_Code/:Page/:Count", apiCurlSend.HallSendGet)
		article.GET("/data/:Site_Code/:Id", apiCurlSend.HallSendGet)
		
		article.POST("", apiServiceSend.CreateArticle)
		
		article.PUT("", apiServiceSend.UpdateArticle)
		
		article.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * image Group
	 */
	image := whaleGroup.Group("/image")
	{
		image.GET("/list", apiCurlSend.HallSendGet)
		image.GET("/page/:Page/:Count", apiCurlSend.HallSendGet)
		
		image.POST("", apiCurlSend.HallSendMultiPartPost)
		image.POST("/update", apiCurlSend.HallSendMultiPartPost)
		
		image.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * home Group
	 */
	home := whaleGroup.Group("/home")
	{
		home.GET("/list/:Start_Time/:End_Time/", apiCurlSend.HallSendGet)
		home.GET("/announcement/:Count/:Page", apiCurlSend.HallSendGet)
		home.GET("/announcementData/:Id", apiCurlSend.HallSendGet)
	}
	
	/**
	 * announcement Group
	 */
	announcement := whaleGroup.Group("/announcement")
	{
		announcement.GET("/list/:Page/:Count", apiCurlSend.HallSendGet)
		announcement.GET("/data/:Id", apiCurlSend.HallSendGet)
		announcement.GET("/hall", apiCurlSend.HallSendGet)
		
		announcement.POST("", apiCurlSend.HallSendPost)
		
		announcement.PUT("", apiCurlSend.HallSendPut)
		
		announcement.DELETE("", apiCurlSend.HallSendDelete)
		announcement.DELETE("/batch", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * bookmark Group
	 */
	bookmark := whaleGroup.Group("/bookmark")
	{
		bookmark.GET("/list", apiCurlSend.SiteSendGet)
		bookmark.GET("/type/:Site_Code", apiCurlSend.HallSendGet)
		bookmark.GET("/backEndList/:Site_Code/:Type_Id", apiCurlSend.HallSendGet)
		bookmark.GET("/data/:Site_Code/:Type_Id/:Id", apiCurlSend.HallSendGet)
		bookmark.GET("/code", apiCurlSend.HallSendGet)
		
		bookmark.POST("", apiCurlSend.HallSendPost)
		bookmark.POST("/type", apiCurlSend.HallSendPost)
		bookmark.POST("/focus", apiCurlSend.SiteSendPost)
		
		bookmark.PUT("/type", apiCurlSend.HallSendPut)
		bookmark.PUT("", apiCurlSend.HallSendPut)
		
		bookmark.DELETE("/focus", apiCurlSend.SiteSendDelete)
		bookmark.DELETE("", apiCurlSend.HallSendDelete)
		bookmark.DELETE("/type", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * bulletin Group
	 */
	bulletin := whaleGroup.Group("/bulletin")
	{
		bulletin.GET("/list", apiCurlSend.SiteSendGet)
		
		bulletin.GET("/backend_list/:Site_Code/:Page/:Count", apiCurlSend.HallSendGet)
		bulletin.GET("/data/:Site_Code/:Id", apiCurlSend.HallSendGet)
		
		bulletin.POST("", apiCurlSend.HallSendPost)
		
		bulletin.PUT("", apiCurlSend.HallSendPut)
		
		bulletin.DELETE("", apiCurlSend.HallSendDelete)
		bulletin.DELETE("/batch", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * page_article Group
	 */
	pageArticle := whaleGroup.Group("/page_article")
	{
		pageArticle.GET("/info", apiCurlSend.SiteSendGet)
		
		pageArticle.GET("/code/:Site_Code", apiCurlSend.HallSendGet)
		pageArticle.GET("/site/:Site_Code", apiCurlSend.HallSendGet)
		
		pageArticle.POST("", apiCurlSend.HallSendPost)
		
		pageArticle.PUT("", apiCurlSend.HallSendPut)
		
		pageArticle.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * message Group
	 */
	message := whaleGroup.Group("/message")
	{
		message.GET("/list", apiCurlSend.SiteSendGet)
		
		message.GET("/backEndList/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		message.GET("/data/:Site_Code/:Id", apiCurlSend.HallSendGet)
		message.GET("/userData/:Site_Code/:Id/:Page/:Count", apiCurlSend.HallSendGet)
		
		message.POST("/read", apiCurlSend.SiteSendPost)
		message.POST("", apiCurlSend.HallSendPost)
	}
	
	/**
	 * customerService Group
	 */
	customerService := whaleGroup.Group("/customer_service")
	{
		customerService.GET("/url", apiCurlSend.SiteSendGet)
	}
}
