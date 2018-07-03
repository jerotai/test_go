package stingray

import (
	"Stingray/helper"
)

/**
 * init Api Post Config
 * apiDto = &dto.Menu{}
 * apiRequestUrl = "menu"
 *
 */
func (i *Stingray) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	default:
		//todo
		helper.HelperLog.ErrorLog("[Stingray InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}