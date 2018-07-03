package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/core/apicurl"
	"Stingray/api/lobster"
)

func InitLobsterRouting(routerGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("lobster_service")
	router := lobster.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)
	
	/**
	 * report Group
	 */
	report := routerGroup.Group("/report")
	{
		report.GET("/outInList/:Site_Code/:Start_Time/:End_Time", apicurlSend.HallSendGet)
		report.GET("/depositList/:Site_Code/:Start_Time/:End_Time/:Page/:Count/", apicurlSend.HallSendGet)
		report.GET("/withdrawList/:Site_Code/:Start_Time/:End_Time/:Page/:Count/", apicurlSend.HallSendGet)
		report.GET("/userTrans/:Site_Code/:Start_Time/:End_Time/:User/:Kind/:Page/:Count/", apicurlSend.HallSendGet)
	}
	
	/**
	 * kind Group
	 */
	kind := routerGroup.Group("/kind")
	{
		kind.GET("/depositType", apicurlSend.HallSendGet)
		kind.GET("/withdrawType", apicurlSend.HallSendGet)
		kind.GET("/transType", apicurlSend.HallSendGet)
	}
}
