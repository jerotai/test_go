package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/api/whitebait"
	"Stingray/core/apicurl"
)


func InitWhitebaitRouting(whitebaitGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("whitebait_service")
	router := whitebait.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)

	
	/**
	 * User Group
	 */
	user := whitebaitGroup.Group("/user")
	{
		user.POST("/login", apicurlSend.WhitebaitLogin)
		
		user.GET("/promo_code/:Promotion_Code/:Promotion_Type", apicurlSend.SiteSendGet)
		user.GET("/passwordWithdrawCheck", apicurlSend.SiteSendGet)
		user.GET("/register", apicurlSend.SiteSendGet)
		user.GET("/plat", apicurlSend.SiteSendGet)
		user.GET("/data", apicurlSend.SiteSendGet)
		user.GET("/balance", apicurlSend.SiteSendGet)
		user.GET("/company_bank", apicurlSend.SiteSendGet)
		user.GET("/transfer", apicurlSend.SiteSendGet)
		user.GET("/alive", apicurlSend.SiteSendGet)
		user.POST("/register", apicurlSend.WhitebaitRegister)
		user.POST("/passwordWithdraw", apicurlSend.SiteSendPost)
		user.PUT("/password", apicurlSend.SiteSendPut)
		user.PUT("/passwordWithdraw", apicurlSend.SiteSendPut)
		
		user.GET("/list/:Site_Code/:Page/:Count/", apicurlSend.HallSendGet)
		user.GET("/info/:id", apicurlSend.HallSendGet)
		user.GET("/registerSetting/:Site_code", apicurlSend.HallSendGet)
		user.POST("/create", apicurlSend.HallSendPost)
		user.PUT("/info", apicurlSend.HallSendPut)
		user.PUT("/registerSetting", apicurlSend.HallSendPut)
		user.PUT("/passwordWithdrawUpdate", apicurlSend.HallSendPut)
		user.PUT("/passwordUpdate", apicurlSend.HallSendPut)
	}
	
	/**
	 * User Level Group
	 */
	userLevel := whitebaitGroup.Group("user_level")
	{
		userLevel.GET("/list/:Site_code", apicurlSend.HallSendGet)
		userLevel.GET("/data/:Id", apicurlSend.HallSendGet)
		userLevel.GET("/amount/:Id/:Type", apicurlSend.HallSendGet)
		userLevel.GET("/payment/:Id", apicurlSend.HallSendGet)
		userLevel.GET("/company_bank/:Id", apicurlSend.HallSendGet)
		
		userLevel.POST("", apicurlSend.HallSendPost)
		
		userLevel.PUT("", apicurlSend.HallSendPut)
		userLevel.PUT("/amount", apicurlSend.HallSendPut)
		userLevel.PUT("/payment", apicurlSend.HallSendPut)
		userLevel.PUT("/company_bank", apicurlSend.HallSendPut)
		userLevel.PUT("/user", apicurlSend.HallSendPut)
		userLevel.PUT("/batch", apicurlSend.HallSendPut)
		
		userLevel.DELETE("", apicurlSend.HallSendDelete)
	}
	
	/**
	 * Bank Group
	 */
	bank := whitebaitGroup.Group("/bank")
	{
		bank.GET("/transfer", apicurlSend.SiteSendGet)
		bank.GET("/third", apicurlSend.SiteSendGet)
	}
}
