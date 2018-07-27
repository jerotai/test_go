package clownfish

import (
	"Stingray/helper"
)


/**
 * init Api Post Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *Clownfish) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	default:
		//todo
		helper.HelperLog.ErrorLog("[Clownfish InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}