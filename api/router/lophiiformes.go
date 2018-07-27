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
	
	var apiCurlSend, _ = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * bank list
	 */
	routerGroup.GET("/transfer_bank/list", apiCurlSend.HallSendGet)
	routerGroup.GET("/third_bank/list", apiCurlSend.HallSendGet)
	
	/**
	 * bank list
	 */
	bankBank := routerGroup.Group("/bank_list")
	{
		bankBank.GET("/fourthThird/:Code", apiCurlSend.HallSendGet)
		bankBank.GET("/fourthTransfer/:Code", apiCurlSend.HallSendGet)
	}
	
	
	/**
	 * company_bank Group
	 */
	companyBank := routerGroup.Group("/company_bank")
	{
		companyBank.GET("/list/:Site_Code", apiCurlSend.HallSendGet)
		companyBank.GET("/data/:Id", apiCurlSend.HallSendGet)
		companyBank.GET("/dropdownList/:Site_Code", apiCurlSend.HallSendGet)
		
		companyBank.POST("", apiCurlSend.HallSendPost)
		
		companyBank.PUT("", apiCurlSend.HallSendPut)
		
		companyBank.DELETE("", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * user_bank Group
	 */
	userBank := routerGroup.Group("/user_bank")
	{
		userBank.GET("/list", apiCurlSend.SiteSendGet)
		userBank.GET("/info/:Id", apiCurlSend.SiteSendGet)
		userBank.POST("", apiCurlSend.SiteSendPost)
		userBank.PUT("", apiCurlSend.SiteSendPut)
		userBank.DELETE("", apiCurlSend.SiteSendDelete)
		
		userBank.GET("/backendList/:User_Id", apiCurlSend.HallSendGet)
		userBank.GET("/backendInfo/:Id/:User_Id", apiCurlSend.HallSendGet)
		userBank.POST("/backend", apiCurlSend.HallSendPost)
		userBank.PUT("/backend", apiCurlSend.HallSendPut)
		userBank.DELETE("/backend", apiCurlSend.HallSendDelete)
	}
	
	/**
	 * payment_config Group
	 */
	paymentConfig := routerGroup.Group("/payment_config")
	{
		paymentConfig.GET("/list/:Site_Code", apiCurlSend.HallSendGet)
		paymentConfig.GET("/default", apiCurlSend.HallSendGet)
		paymentConfig.GET("/info/:Id", apiCurlSend.HallSendGet)
		paymentConfig.GET("/banks/:Fourth_Id", apiCurlSend.HallSendGet)
		
		paymentConfig.POST("", apiCurlSend.HallSendPost)
		
		paymentConfig.PUT("", apiCurlSend.HallSendPut)
		paymentConfig.PUT("/banks", apiCurlSend.HallSendPut)
		
		paymentConfig.DELETE("", apiCurlSend.HallSendDelete)
	}
}
