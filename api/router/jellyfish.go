package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/core/apicurl"
	"Stingray/api/jellyfish"
)

func InitJellyFishRouting(jellyfishGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("jellyfish_service")
	router := jellyfish.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	var apiCurlSend = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * User Group
	 */
	jellyfishGroup.POST(("/login"), apiCurlSend.JellyFishLogin)
	
	/**
	 * Password Group
	 */
	jellyfishGroup.PUT(("/password"), apiCurlSend.HallSendPut)
	
	/**
	 * Role Group
	 */
	jellyfishGroup.GET(("/role/list"), apiCurlSend.HallSendGet)

	/**
	 * Hall Group (廳主)
	 */
	hall := jellyfishGroup.Group("/hall")
	{
		hall.POST((""), apiCurlSend.HallSendPost)
		hall.PUT((""), apiCurlSend.HallSendPut)
		hall.PUT(("/enabled"), apiCurlSend.HallSendPut)
		hall.PUT(("/password"), apiCurlSend.HallSendPut)
		hall.GET(("/subList/:Page/:Count"), apiCurlSend.HallSendGet)
		hall.GET(("/dropdownlist"), apiCurlSend.HallSendGet)
	}
	
	/**
	 * Shareholder Group (股東)
	 */
	shareholder := jellyfishGroup.Group("/shareholder")
	{
		//hall_code=CQ1&site_code=AA01&status=1&name=AA01
		//&start_time=2018-05-01 00:00:00&end_time=2018-05-12 23:59:59
		shareholder.GET(("/list/:Site_Code/:Page/:Count/"), apiCurlSend.HallSendGet)
		shareholder.GET(("/data/:Id"), apiCurlSend.HallSendGet)
		shareholder.GET(("/dropdownlist/:Site_Code"), apiCurlSend.HallSendGet)
		shareholder.POST((""), apiCurlSend.HallSendPost)
		shareholder.PUT((""), apiCurlSend.HallSendPut)
		shareholder.PUT(("/password"), apiCurlSend.HallSendPut)
	}
	
	/**
	 * ShareholderSub Group (股東子帳號)
	 */
	shareholderSub := jellyfishGroup.Group("/shareholderSub")
	{
		shareholderSub.GET(("/list/:Page/:Count/"), apiCurlSend.HallSendGet)
		shareholderSub.GET(("/data/:Id"), apiCurlSend.HallSendGet)
		shareholderSub.POST((""), apiCurlSend.HallSendPost)
		shareholderSub.PUT((""), apiCurlSend.HallSendPut)
		shareholderSub.PUT(("/password"), apiCurlSend.HallSendPut)
	}
	
	/**
	 * hallSub Grpup (廳主子帳號)
	 */
	hallSub := jellyfishGroup.Group("/hallSub")
	{
		hallSub.POST((""), apiCurlSend.HallSendPost)
		hallSub.PUT("", apiCurlSend.HallSendPut)
		hallSub.PUT(("/password"), apiCurlSend.HallSendPut)
		hallSub.GET("/list/:Page/:Count", apiCurlSend.HallSendGet)
		hallSub.GET("/data/:Id", apiCurlSend.HallSendGet)
	}
	
	/**
	 * GeneralAgent Grpup (總代理管理)
	 */
	generalAgent := jellyfishGroup.Group("/generalAgent")
	{
		generalAgent.GET("/list/:Shareholder_Id/:Page/:Count/", apiCurlSend.HallSendGet)
		generalAgent.GET("/data/:Id", apiCurlSend.HallSendGet)
		generalAgent.GET(("/dropdownlist/:Site_Code"), apiCurlSend.HallSendGet)
		generalAgent.POST("", apiCurlSend.HallSendPost)
		generalAgent.PUT("", apiCurlSend.HallSendPut)
		generalAgent.PUT("/password", apiCurlSend.HallSendPut)
	}
	
	/**
	 * GeneralAgentSub Grpup (總代理子帳號管理	)
	 */
	generalAgentSub := jellyfishGroup.Group("/generalAgentSub")
	{
		generalAgentSub.GET("/list/:Page/:Count/", apiCurlSend.HallSendGet)
		generalAgentSub.GET("/data/:Id", apiCurlSend.HallSendGet)
		generalAgentSub.POST("", apiCurlSend.HallSendPost)
		generalAgentSub.PUT("", apiCurlSend.HallSendPut)
		generalAgentSub.PUT("/password", apiCurlSend.HallSendPut)
	}
	
	/**
	 * Agent Grpup (代理管理)
	 */
	agent := jellyfishGroup.Group("/agent")
	{
		agent.GET("/list/:GeneralAgent_Id/:Page/:Count/", apiCurlSend.HallSendGet)
		agent.GET("/data/:Id", apiCurlSend.HallSendGet)
		agent.GET(("/dropdownlist/:Site_Code"), apiCurlSend.HallSendGet)
		agent.POST("", apiCurlSend.HallSendPost)
		agent.PUT("", apiCurlSend.HallSendPut)
		agent.PUT("/password", apiCurlSend.HallSendPut)
	}
	
	/**
	 * AgentSub Grpup (代理子帳號管理)
	 */
	agentSub := jellyfishGroup.Group("/agentSub")
	{
		agentSub.GET("/list/:Page/:Count/", apiCurlSend.HallSendGet)
		agentSub.GET("/data/:Id", apiCurlSend.HallSendGet)
		agentSub.POST("", apiCurlSend.HallSendPost)
		agentSub.PUT("", apiCurlSend.HallSendPut)
		agentSub.PUT("/password", apiCurlSend.HallSendPut)
	}
	
	
	/**
	 * authGroup Grpup (權限群組)
	 */
	authgroup := jellyfishGroup.Group("authGroup")
	{
		authgroup.POST("", apiCurlSend.HallSendPost)
		authgroup.PUT((""), apiCurlSend.HallSendPut)
		authgroup.GET("/list/:Page/:Count", apiCurlSend.HallSendGet)
		authgroup.GET("/detail/:Group_Id", apiCurlSend.HallSendGet)
		authgroup.GET("/dropDown", apiCurlSend.HallSendGet)
	}
	
	/**
	 * agentSystem Grpup (體系)
	 */
	agentSystem := jellyfishGroup.Group("/agentSystem")
	{
		agentSystem.GET("/siteTotal/:Site_Code/:Page/:Count", apiCurlSend.HallSendGet)
		agentSystem.GET("/shareholderTotal/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		agentSystem.GET("/generalAgentTotal/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		agentSystem.GET("/agentTotal/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		agentSystem.GET("/memberTotal/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
	}
}
