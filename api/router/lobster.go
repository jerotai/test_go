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
	
	var apiCurlSend, _ = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * report Group
	 */
	report := routerGroup.Group("/report")
	{
		report.GET("/outInList/:Site_Code/:Start_Time/:End_Time/:Order", apiCurlSend.HallSendGet)
		report.GET("/depositList/:Site_Code/:Start_Time/:End_Time/:Page/:Count/", apiCurlSend.HallSendGet)
		report.GET("/withdrawList/:Site_Code/:Start_Time/:End_Time/:Page/:Count/", apiCurlSend.HallSendGet)
		report.GET("/userTrans/:Site_Code/:Start_Time/:End_Time/:User/:Kind/:Page/:Count/:Order/", apiCurlSend.HallSendGet)
	}
	
	/**
	 * kind Group
	 */
	kind := routerGroup.Group("/kind")
	{
		kind.GET("/depositType", apiCurlSend.HallSendGet)
		kind.GET("/withdrawType", apiCurlSend.HallSendGet)
		kind.GET("/transType", apiCurlSend.HallSendGet)
	}
	
	/**
	 * userLoginReport Group
	 */
	userLoginReport := routerGroup.Group("/userLoginReport")
	{
		userLoginReport.GET("/userLoginDriveCount/:Site_Code/:Start_Time/:End_Time",apiCurlSend.HallSendGet)
		userLoginReport.GET("/userLoginInfo/:Site_Code/:Start_Time/:End_Time/:Ip_Is_Repeat/:Page/:Count/",apiCurlSend.HallSendGet)
		userLoginReport.GET("/userLoginInfoOnline/:Site_Code/:Start_Time/:End_Time/:Ip_Is_Repeat/:Page/:Count/",apiCurlSend.HallSendGet)
		userLoginReport.GET("/userLoginRecord/:Site_Code/:Start_Time/:End_Time/:Page/:Count/",apiCurlSend.HallSendGet)
	}
	
	betting := routerGroup.Group("betting")
	{
		betting.GET("/subordinate/:Site_Code/:Start_Time/:End_Time/", apiCurlSend.HallSendGet)
		betting.GET("/subordinateSummary/:Site_Code/:Start_Time/:End_Time", apiCurlSend.HallSendGet)
	}
}
