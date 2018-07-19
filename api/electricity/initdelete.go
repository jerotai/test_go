package electricity

import (
	"Stingray/helper"
)


/**
 * init Api Delete Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *Electricity) InitDeleteApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	default:
		//todo
		helper.HelperLog.ErrorLog("[Electricity InitDeleteApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}