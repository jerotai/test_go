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
	
	var apiCurlSend, apiServiceSend = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * User Group
	 */
	user := whitebaitGroup.Group("/user")
	{
		user.POST("/login", apiServiceSend.WhitebaitLogin)
		
		user.GET("/promo_code/:Promotion_Code/:Promotion_Type", apiCurlSend.SiteSendGet)
		user.GET("/passwordWithdrawCheck", apiCurlSend.SiteSendGet)
		user.GET("/register", apiCurlSend.SiteSendGet)
		user.GET("/plat", apiCurlSend.SiteSendGet)
		user.GET("/data", apiCurlSend.SiteSendGet)
		user.GET("/balance", apiCurlSend.SiteSendGet)
		user.GET("/company_bank", apiCurlSend.SiteSendGet)
		user.GET("/transfer", apiCurlSend.SiteSendGet)
		user.GET("/alive", apiCurlSend.SiteSendGet)
		user.POST("/register", apiServiceSend.WhitebaitRegister)
		user.POST("/passwordWithdraw", apiCurlSend.SiteSendPost)
		user.PUT("/password", apiCurlSend.SiteSendPut)
		user.PUT("/passwordWithdraw", apiCurlSend.SiteSendPut)
		
		user.GET("/list/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		user.GET("/blurry/:Site_Code/:Account", apiCurlSend.HallSendGet)
		user.GET("/info/:id", apiCurlSend.HallSendGet)
		user.GET("/registerSetting/:Site_code", apiCurlSend.HallSendGet)
		user.POST("/create", apiCurlSend.HallSendPost)
		user.PUT("/info", apiCurlSend.HallSendPut)
		user.PUT("/registerSetting", apiCurlSend.HallSendPut)
		user.PUT("/passwordWithdrawUpdate", apiCurlSend.HallSendPut)
		user.PUT("/passwordUpdate", apiCurlSend.HallSendPut)
	}
	
	/**
	 * User Level Group
	 */
	userLevel := whitebaitGroup.Group("user_level")
	{
		userLevel.GET("/list/:Site_code/:Page/:Count/", apiCurlSend.HallSendGet)
		userLevel.GET("/data/:Id", apiCurlSend.HallSendGet)
		userLevel.GET("/amount/:Id/:Type", apiCurlSend.HallSendGet)
		userLevel.GET("/payment/:Id", apiCurlSend.HallSendGet)
		userLevel.GET("/company_bank/:Id", apiCurlSend.HallSendGet)
		userLevel.GET("/user_list/:Id/:Page/:Count", apiCurlSend.HallSendGet)
		userLevel.GET("/dropdownList/:Site_code", apiCurlSend.HallSendGet)
		userLevel.GET("/userPreview/:User_Ids/:User_Level_Id", apiCurlSend.HallSendGet)
		userLevel.GET("/batchPreview/:Before_User_Level_Id/:After_User_Level_Id", apiCurlSend.HallSendGet)
		
		userLevel.POST("", apiCurlSend.HallSendPost)
		
		userLevel.PUT("", apiCurlSend.HallSendPut)
		userLevel.PUT("/amount", apiCurlSend.HallSendPut)
		userLevel.PUT("/payment", apiCurlSend.HallSendPut)
		userLevel.PUT("/company_bank", apiCurlSend.HallSendPut)
		userLevel.PUT("/user", apiCurlSend.HallSendPut)
		userLevel.PUT("/batch", apiCurlSend.HallSendPut)
		
		userLevel.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * Bank Group
	 */
	bank := whitebaitGroup.Group("/bank")
	{
		bank.GET("/transfer", apiCurlSend.SiteSendGet)
		bank.GET("/third", apiCurlSend.SiteSendGet)
	}
	
	/**
	 * guest Group
	 */
	guest := whitebaitGroup.Group("/guest")
	{
		guest.GET("/myself", apiCurlSend.HallSendGet)
		guest.GET("/hall", apiCurlSend.HallSendGet)
		guest.GET("/site", apiCurlSend.HallSendGet)
		
		guest.POST("/login", apiServiceSend.WhitebaitGuestLogin)
		
		guest.PUT("/myself", apiCurlSend.HallSendPut)
		guest.PUT("/hall", apiCurlSend.HallSendPut)
		guest.PUT("/site", apiCurlSend.HallSendPut)
	}
}
