package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/api/lophiiformes"
	"Stingray/core/apicurl"
)

func InitLophiiformesRouting(routerGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("lophiiformes_service")
	router := lophiiformes.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)
	
	/**
	 * bank list
	 */
	routerGroup.GET("/transfer_bank/list", apicurlSend.HallSendGet)
	routerGroup.GET("/third_bank/list", apicurlSend.HallSendGet)
	
	/**
	 * company_bank Group
	 */
	companyBank := routerGroup.Group("/company_bank")
	{
		companyBank.GET("/list/:Site_Code", apicurlSend.HallSendGet)
		companyBank.GET("/data/:Id", apicurlSend.HallSendGet)
		companyBank.GET("/dropdownList/:Site_Code", apicurlSend.HallSendGet)
		
		companyBank.POST("", apicurlSend.HallSendPost)
		
		companyBank.PUT("", apicurlSend.HallSendPut)
		
		companyBank.DELETE("", apicurlSend.HallSendDelete)
	}
	
	/**
	 * user_bank Group
	 */
	userBank := routerGroup.Group("/user_bank")
	{
		userBank.GET("/list", apicurlSend.SiteSendGet)
		userBank.GET("/info/:Id", apicurlSend.SiteSendGet)
		userBank.POST("", apicurlSend.SiteSendPost)
		userBank.PUT("", apicurlSend.SiteSendPut)
		userBank.DELETE("", apicurlSend.SiteSendDelete)
		
		userBank.GET("/backendList/:User_Id", apicurlSend.HallSendGet)
		userBank.GET("/backendInfo/:Id/:User_Id", apicurlSend.HallSendGet)
		userBank.POST("/backend", apicurlSend.HallSendPost)
		userBank.PUT("/backend", apicurlSend.HallSendPut)
		userBank.DELETE("/backend", apicurlSend.HallSendDelete)
	}
	
	/**
	 * payment_config Group
	 */
	paymentConfig := routerGroup.Group("/payment_config")
	{
		paymentConfig.GET("/list", apicurlSend.HallSendGet)
		paymentConfig.GET("/default", apicurlSend.HallSendGet)
		paymentConfig.GET("/info/:Id", apicurlSend.HallSendGet)
		paymentConfig.GET("/banks/:Fourth_Id", apicurlSend.HallSendGet)
		
		paymentConfig.POST("", apicurlSend.HallSendPost)
		
		paymentConfig.PUT("", apicurlSend.HallSendPut)
		paymentConfig.PUT("/banks", apicurlSend.HallSendPut)
		
		paymentConfig.DELETE("", apicurlSend.HallSendDelete)
	}
}
