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
	
	apicurlSend := apicurl.New(apiConf, apiConfInit)
	
	
	/**
	 * User Group
	 */
	jellyfishGroup.POST(("/login"), apicurlSend.JellyFishLogin)
	
	/**
	 * Password Group
	 */
	jellyfishGroup.PUT(("/password"), apicurlSend.HallSendPut)
	
	/**
	 * Role Group
	 */
	jellyfishGroup.GET(("/role/list"), apicurlSend.HallSendGet)

	/**
	 * Hall Group (廳主)
	 */
	hall := jellyfishGroup.Group("/hall")
	{
		hall.POST((""), apicurlSend.HallSendPost)
		hall.PUT((""), apicurlSend.HallSendPut)
		hall.PUT(("/enabled"), apicurlSend.HallSendPut)
		hall.PUT(("/password"), apicurlSend.HallSendPut)
		hall.GET(("/subList/:Page/:Count"), apicurlSend.HallSendGet)
		hall.GET(("/dropdownlist"), apicurlSend.HallSendGet)
	}
	
	/**
	 * Shareholder Group (股東)
	 */
	shareholder := jellyfishGroup.Group("/shareholder")
	{
		//hall_code=CQ1&site_code=AA01&status=1&name=AA01
		//&start_time=2018-05-01 00:00:00&end_time=2018-05-12 23:59:59
		shareholder.GET(("/list/:Site_Code/:Page/:Count/"), apicurlSend.HallSendGet)
		shareholder.GET(("/data/:Id"), apicurlSend.HallSendGet)
		shareholder.GET(("/dropdownlist/:Site_Code"), apicurlSend.HallSendGet)
		shareholder.POST((""), apicurlSend.HallSendPost)
		shareholder.PUT((""), apicurlSend.HallSendPut)
		shareholder.PUT(("/password"), apicurlSend.HallSendPut)
	}
	
	/**
	 * ShareholderSub Group (股東子帳號)
	 */
	shareholderSub := jellyfishGroup.Group("/shareholderSub")
	{
		shareholderSub.GET(("/list/:Page/:Count/"), apicurlSend.HallSendGet)
		shareholderSub.GET(("/data/:Id"), apicurlSend.HallSendGet)
		shareholderSub.POST((""), apicurlSend.HallSendPost)
		shareholderSub.PUT((""), apicurlSend.HallSendPut)
		shareholderSub.PUT(("/password"), apicurlSend.HallSendPut)
	}
	
	/**
	 * hallSub Grpup (廳主子帳號)
	 */
	hallSub := jellyfishGroup.Group("/hallSub")
	{
		hallSub.POST((""), apicurlSend.HallSendPost)
		hallSub.PUT("", apicurlSend.HallSendPut)
		hallSub.PUT(("/password"), apicurlSend.HallSendPut)
		hallSub.GET("/list/:Page/:Count",apicurlSend.HallSendGet)
		hallSub.GET("/data/:Id",apicurlSend.HallSendGet)
		
	}
	
	/**
	 * GeneralAgent Grpup (總代理管理)
	 */
	generalAgent := jellyfishGroup.Group("/generalAgent")
	{
		generalAgent.GET("/list/:Shareholder_Id/:Page/:Count/",apicurlSend.HallSendGet)
		generalAgent.GET("/data/:Id",apicurlSend.HallSendGet)
		generalAgent.GET(("/dropdownlist/:Site_Code"), apicurlSend.HallSendGet)
		generalAgent.POST("", apicurlSend.HallSendPost)
		generalAgent.PUT("", apicurlSend.HallSendPut)
		generalAgent.PUT("/password", apicurlSend.HallSendPut)
	}
	
	/**
	 * GeneralAgentSub Grpup (總代理子帳號管理	)
	 */
	generalAgentSub := jellyfishGroup.Group("/generalAgentSub")
	{
		generalAgentSub.GET("/list/:Page/:Count/",apicurlSend.HallSendGet)
		generalAgentSub.GET("/data/:Id",apicurlSend.HallSendGet)
		generalAgentSub.POST("", apicurlSend.HallSendPost)
		generalAgentSub.PUT("", apicurlSend.HallSendPut)
		generalAgentSub.PUT("/password", apicurlSend.HallSendPut)
	}
	
	/**
	 * Agent Grpup (代理管理)
	 */
	agent := jellyfishGroup.Group("/agent")
	{
		agent.GET("/list/:GeneralAgent_Id/:Page/:Count/",apicurlSend.HallSendGet)
		agent.GET("/data/:Id",apicurlSend.HallSendGet)
		agent.GET(("/dropdownlist/:Site_Code"), apicurlSend.HallSendGet)
		agent.POST("", apicurlSend.HallSendPost)
		agent.PUT("", apicurlSend.HallSendPut)
		agent.PUT("/password", apicurlSend.HallSendPut)
	}
	
	/**
	 * AgentSub Grpup (代理子帳號管理)
	 */
	agentSub := jellyfishGroup.Group("/agentSub")
	{
		agentSub.GET("/list/:Page/:Count/",apicurlSend.HallSendGet)
		agentSub.GET("/data/:Id",apicurlSend.HallSendGet)
		agentSub.POST("", apicurlSend.HallSendPost)
		agentSub.PUT("", apicurlSend.HallSendPut)
		agentSub.PUT("/password", apicurlSend.HallSendPut)
	}
	
	
	/**
	 * authGroup Grpup (權限群組)
	 */
	authgroup := jellyfishGroup.Group("authGroup")
	{
		authgroup.POST("", apicurlSend.HallSendPost)
		authgroup.PUT((""), apicurlSend.HallSendPut)
		authgroup.GET("/list/:Page/:Count",apicurlSend.HallSendGet)
		authgroup.GET("/detail/:Group_Id",apicurlSend.HallSendGet)
		authgroup.GET("/dropDown",apicurlSend.HallSendGet)
	}
	
	/**
	 * agentSystem Grpup (體系)
	 */
	agentSystem := jellyfishGroup.Group("/agentSystem")
	{
		agentSystem.GET("/siteTotal/:Site_Code/:Page/:Count",apicurlSend.HallSendGet)
		agentSystem.GET("/shareholderTotal/:Site_Code/:Page/:Count/",apicurlSend.HallSendGet)
		agentSystem.GET("/generalAgentTotal/:Site_Code/:Page/:Count/",apicurlSend.HallSendGet)
		agentSystem.GET("/agentTotal/:Site_Code/:Page/:Count/",apicurlSend.HallSendGet)
		agentSystem.GET("/memberTotal/:Site_Code/:Page/:Count/",apicurlSend.HallSendGet)
	}
}
