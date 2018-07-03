package lobster

import (
	"Stingray/helper"
)

/**
 * init Api Post Config
 * apiDto = &dto.Menu{}
 * apiRequestUrl = "menu"
 *
 */
func (i *Lobster) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lobster InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}