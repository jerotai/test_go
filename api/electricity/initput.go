package electricity

import (
	"Stingray/helper"
	dto "Stingray/core/dto/electricitydto"
)


/**
 * init Api Put Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *Electricity) InitPutApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/reward_config":
		apiDto = &dto.UpdateRewardConfig{}
		apiRequestUrl = "reward_config"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Electricity InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
