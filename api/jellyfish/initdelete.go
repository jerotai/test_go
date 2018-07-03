package jellyfish

import (
	"Stingray/helper"
)

/**
 * init Api Delete Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *JellyFish) InitDeleteApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	
	switch apiUrl {
	default:
		//todo
		helper.HelperLog.ErrorLog("[JellyFish InitDeleteApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}