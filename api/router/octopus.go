package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/api/octopus"
	"Stingray/core/apicurl"
)


func InitOctopusRouting(routerGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("octopus_service")
	router := octopus.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)
	
	/**
	 * manualDeposits Group
	 */
	manualDeposits := routerGroup.Group("/manualDeposits")
	{
		manualDeposits.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apicurlSend.HallSendGet)
		manualDeposits.GET("/audit/:Txnid", apicurlSend.HallSendGet)
		
		manualDeposits.POST("", apicurlSend.HallSendPost)
		manualDeposits.POST("/pass", apicurlSend.HallSendPost)
		manualDeposits.POST("/reject", apicurlSend.HallSendPost)
		
		manualDeposits.PUT("/unlock", apicurlSend.HallSendPut)
	}
	
	/**
	 * manualWithdraws Group 人工出款
	 */
	manualWithdraws := routerGroup.Group("/manualWithdraws")
	{
		manualWithdraws.POST("", apicurlSend.HallSendPost)
		manualWithdraws.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apicurlSend.HallSendGet)
		
		manualWithdraws.PUT("/unlock", apicurlSend.HallSendPut)
	}

	/**
	 * fourthDeposits Group
	 */
	fourthDeposits := routerGroup.Group("/fourthDeposits")
	{
		fourthDeposits.GET("/list/:Site_code/:Status/:Start_Time/:End_Time/:Page/:Count/", apicurlSend.HallSendGet)
		fourthDeposits.GET("/menu/:Site_code", apicurlSend.HallSendGet)
		fourthDeposits.GET("/audit/:Txnid", apicurlSend.HallSendGet)
		
		fourthDeposits.GET("/limit", apicurlSend.SiteSendGet)
		fourthDeposits.GET("/menu_thirds", apicurlSend.SiteSendGet)
		fourthDeposits.GET("/menu_fourths/:Bank_Id", apicurlSend.SiteSendGet)
		
		fourthDeposits.POST("", apicurlSend.SiteSendPost)
		
		fourthDeposits.POST("/pass", apicurlSend.HallSendPost)
		fourthDeposits.POST("/reject", apicurlSend.HallSendPost)
		
		fourthDeposits.PUT("/unlock", apicurlSend.HallSendPut)
	}

	/**
	 * bankDeposits Group
	 */
	bankDeposits := routerGroup.Group("/bankDeposits")
	{
		bankDeposits.GET("/menu/:Site_Code/:Status/:Start_Time/:End_Time/:Page/:Count/", apicurlSend.HallSendGet)
		bankDeposits.GET("/front_menu", apicurlSend.SiteSendGet)
		
		bankDeposits.GET("/list/:Site_Code/:Status/:Start_Time/:End_Time/:Page/:Count/", apicurlSend.HallSendGet)
		bankDeposits.GET("/audit/:Txnid", apicurlSend.HallSendGet)
		bankDeposits.GET("/menu/:Site_Code", apicurlSend.HallSendGet)
		
		bankDeposits.POST("", apicurlSend.SiteSendPost)
		
		bankDeposits.POST("/pass", apicurlSend.HallSendPost)
		bankDeposits.POST("/reject", apicurlSend.HallSendPost)
		
		bankDeposits.PUT("/unlock", apicurlSend.HallSendPut)
	}

	/**
	 * providerDeposits Group
	 */
	providerDeposits := routerGroup.Group("/providerDeposits")
	{
		providerDeposits.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apicurlSend.HallSendGet)
		providerDeposits.GET("/audit/:Txnid", apicurlSend.HallSendGet)
		providerDeposits.GET("/limit", apicurlSend.SiteSendGet)
		
		providerDeposits.POST("/pass", apicurlSend.HallSendPost)
		providerDeposits.POST("/reject", apicurlSend.HallSendPost)
		providerDeposits.POST("", apicurlSend.SiteSendPost)
		
		providerDeposits.PUT("/unlock", apicurlSend.HallSendPut)
	}
	
	/**
	 * providerWithdraws Group (轉帳出款)
	 */
	providerWithdraws := routerGroup.Group("/providerWithdraws")
	{
		providerWithdraws.GET("/setting", apicurlSend.SiteSendGet)
		providerWithdraws.GET("/limit", apicurlSend.SiteSendGet)
		
		providerWithdraws.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apicurlSend.HallSendGet)
		
		providerWithdraws.POST("", apicurlSend.SiteSendPost)
		
		providerWithdraws.PUT("/unlock", apicurlSend.HallSendPut)
	}
	
	/**
	 * trans Group
	 */
	trans := routerGroup.Group("/trans")
	{
		trans.PUT("/unlock", apicurlSend.HallSendPut)
		trans.PUT("/selfunlock", apicurlSend.HallSendPut)
	}
	
	/**
	 * transReport Group
	 */
	asi := routerGroup.Group("/transReports")
	{
		asi.GET("/list/:Start_Time/:End_Time", apicurlSend.HallSendGet)
		asi.GET("/deposit", apicurlSend.SiteSendGet)
		asi.GET("/withdrawList", apicurlSend.SiteSendGet)
	}
	
	/**
	 * transSumReport Group
	 */
	tsr := routerGroup.Group("/transSumReport")
	{
		tsr.GET("/depositList/:Start_Time/:End_Time", apicurlSend.HallSendGet)
		tsr.GET("/withdrawList/:Start_Time/:End_Time", apicurlSend.HallSendGet)
	}
}
