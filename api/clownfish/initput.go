package clownfish

import (
	"Stingray/helper"
)


/**
 * init Api Put Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *Clownfish) InitPutApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	default:
		//todo
		helper.HelperLog.ErrorLog("[Clownfish InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
