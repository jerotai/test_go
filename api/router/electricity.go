package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/helper"
	"Stingray/core/apicurl"
	"Stingray/api/electricity"
)

func InitElectricityRouting(electricityGroup *gin.RouterGroup) {
	apiConf := helper.ApiSetting("electricity_service")
	router := electricity.New()
	
	apiConfInit := &apicurl.ApiConfInit{}
	apiConfInit.InitGetApiConfig = router.InitGetApiConfig
	apiConfInit.InitPostApiConfig = router.InitPostApiConfig
	apiConfInit.InitPutApiConfig = router.InitPutApiConfig
	apiConfInit.InitDeleteApiConfig = router.InitDeleteApiConfig
	
	var apiCurlSend, _ = apicurl.GetCurlSend(apiConf, apiConfInit)
	
	/**
	 * rewardConfig Group
	 */
	rewardConfig := electricityGroup.Group("/reward_config")
	{
		rewardConfig.GET("/setting/:Site_Code", apiCurlSend.HallSendGet)
		
		rewardConfig.PUT("", apiCurlSend.HallSendPut)
	}
	
	/**
	 * reward Group
	 */
	reward := electricityGroup.Group("/reward")
	{
		reward.GET("/info", apiCurlSend.SiteSendGet)
		reward.GET("/record/:Issue", apiCurlSend.SiteSendGet)
		reward.GET("/dropdown", apiCurlSend.SiteSendGet)
		
		reward.GET("/list/:Site_Code/:Status/:Count/:Page/", apiCurlSend.HallSendGet)
		reward.GET("/dropdownList/:Site_Code", apiCurlSend.HallSendGet)
		reward.GET("/reject/:Items/:Remark", apiCurlSend.HallSendGet)
		reward.GET("/pass/:Items/:Dama_Ratio/:Remark", apiCurlSend.HallSendGet)
		reward.GET("/report/:Site_Code/:Page/:Count/", apiCurlSend.HallSendGet)
		reward.GET("/reportInvitee/:Site_Code/:Referrer_Id/:Page/:Count/", apiCurlSend.HallSendGet)
	}
}