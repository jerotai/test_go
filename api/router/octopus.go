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
	
	var apiCurlSend, _ = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * manualDeposits Group
	 */
	manualDeposits := routerGroup.Group("/manualDeposits")
	{
		manualDeposits.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		manualDeposits.GET("/audit/:Txnid", apiCurlSend.HallSendGet)
		
		manualDeposits.POST("", apiCurlSend.HallSendPost)
		manualDeposits.POST("/pass", apiCurlSend.HallSendPost)
		manualDeposits.POST("/reject", apiCurlSend.HallSendPost)
		
		manualDeposits.PUT("/unlock", apiCurlSend.HallSendPut)
	}
	
	/**
	 * manualWithdraws Group 人工出款
	 */
	manualWithdraws := routerGroup.Group("/manualWithdraws")
	{
		manualWithdraws.POST("", apiCurlSend.HallSendPost)
		manualWithdraws.POST("/auditPass", apiCurlSend.HallSendPost)
		manualWithdraws.POST("/auditReject", apiCurlSend.HallSendPost)
		manualWithdraws.POST("/grantPass", apiCurlSend.HallSendPost)
		manualWithdraws.POST("/grantReject", apiCurlSend.HallSendPost)
		
		manualWithdraws.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		manualWithdraws.GET("/audit/:Txnid/:Type", apiCurlSend.HallSendGet)
		manualWithdraws.GET("/grant/:Txnid/:Type", apiCurlSend.HallSendGet)
		
		manualWithdraws.PUT("/unlock", apiCurlSend.HallSendPut)
	}

	/**
	 * fourthDeposits Group
	 */
	fourthDeposits := routerGroup.Group("/fourthDeposits")
	{
		fourthDeposits.GET("/list/:Site_code/:Status/:Start_Time/:End_Time/:Page/:Count/", apiCurlSend.HallSendGet)
		fourthDeposits.GET("/menu/:Site_code", apiCurlSend.HallSendGet)
		fourthDeposits.GET("/audit/:Txnid", apiCurlSend.HallSendGet)
		
		fourthDeposits.GET("/limit", apiCurlSend.SiteSendGet)
		fourthDeposits.GET("/menu_thirds", apiCurlSend.SiteSendGet)
		fourthDeposits.GET("/menu_fourths/:Bank_Id", apiCurlSend.SiteSendGet)
		
		fourthDeposits.POST("", apiCurlSend.SiteSendPost)
		
		fourthDeposits.POST("/pass", apiCurlSend.HallSendPost)
		fourthDeposits.POST("/reject", apiCurlSend.HallSendPost)
		
		fourthDeposits.PUT("/unlock", apiCurlSend.HallSendPut)
	}

	/**
	 * bankDeposits Group
	 */
	bankDeposits := routerGroup.Group("/bankDeposits")
	{
		bankDeposits.GET("/menu/:Site_Code/:Status/:Start_Time/:End_Time/:Page/:Count/", apiCurlSend.HallSendGet)
		bankDeposits.GET("/front_menu", apiCurlSend.SiteSendGet)
		
		bankDeposits.GET("/list/:Site_Code/:Status/:Start_Time/:End_Time/:Page/:Count/", apiCurlSend.HallSendGet)
		bankDeposits.GET("/audit/:Txnid", apiCurlSend.HallSendGet)
		bankDeposits.GET("/menu/:Site_Code", apiCurlSend.HallSendGet)
		
		bankDeposits.POST("", apiCurlSend.SiteSendPost)
		
		bankDeposits.POST("/pass", apiCurlSend.HallSendPost)
		bankDeposits.POST("/reject", apiCurlSend.HallSendPost)
		
		bankDeposits.PUT("/unlock", apiCurlSend.HallSendPut)
	}

	/**
	 * providerDeposits Group
	 */
	providerDeposits := routerGroup.Group("/providerDeposits")
	{
		providerDeposits.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		providerDeposits.GET("/audit/:Txnid", apiCurlSend.HallSendGet)
		providerDeposits.GET("/limit", apiCurlSend.SiteSendGet)
		
		providerDeposits.POST("/pass", apiCurlSend.HallSendPost)
		providerDeposits.POST("/reject", apiCurlSend.HallSendPost)
		providerDeposits.POST("", apiCurlSend.SiteSendPost)
		
		providerDeposits.PUT("/unlock", apiCurlSend.HallSendPut)
	}
	
	/**
	 * providerWithdraws Group (轉帳出款)
	 */
	providerWithdraws := routerGroup.Group("/providerWithdraws")
	{
		providerWithdraws.GET("/limit", apiCurlSend.SiteSendGet)
		
		providerWithdraws.GET("/list/:Status/:Start_Time/:End_Time/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		providerWithdraws.GET("/audit/:Txnid/:Type", apiCurlSend.HallSendGet)
		providerWithdraws.GET("/grant/:Txnid/:Type", apiCurlSend.HallSendGet)
		
		providerWithdraws.POST("", apiCurlSend.SiteSendPost)
		providerWithdraws.POST("/auditPass", apiCurlSend.HallSendPost)
		providerWithdraws.POST("/auditReject", apiCurlSend.HallSendPost)
		providerWithdraws.POST("/grantPass", apiCurlSend.HallSendPost)
		providerWithdraws.POST("/grantReject", apiCurlSend.HallSendPost)
		
		providerWithdraws.PUT("/unlock", apiCurlSend.HallSendPut)
	}
	
	/**
	 * trans Group
	 */
	trans := routerGroup.Group("/trans")
	{
		trans.PUT("/unlock", apiCurlSend.HallSendPut)
		trans.PUT("/selfunlock", apiCurlSend.HallSendPut)
	}
	
	/**
	 * transReport Group
	 */
	asi := routerGroup.Group("/transReports")
	{
		asi.GET("/list/:Start_Time/:End_Time", apiCurlSend.HallSendGet)
		asi.GET("/deposit", apiCurlSend.SiteSendGet)
		asi.GET("/withdrawList", apiCurlSend.SiteSendGet)
		asi.GET("/moneyDetail/:Category/:Option/:Start_Time/:End_Time", apiCurlSend.SiteSendGet)
	}
	
	/**
	 * transSumReport Group
	 */
	tsr := routerGroup.Group("/transSumReport")
	{
		tsr.GET("/depositList/:Start_Time/:End_Time/:Site_Code", apiCurlSend.HallSendGet)
		tsr.GET("/withdrawList/:Start_Time/:End_Time/:Site_Code", apiCurlSend.HallSendGet)
	}
}
