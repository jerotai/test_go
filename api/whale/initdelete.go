package whale

import (
	dto "Stingray/core/dto/whaledto"
	"Stingray/helper"
)

/**
 * init Api Delete Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Whale) InitDeleteApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/banner":
		apiDto = &dto.DeleteBanner{}
		apiRequestUrl = "banner"
	case "/promotion":
		apiDto = &dto.DeletePromotion{}
		apiRequestUrl = "promotion"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whale InitDeleteApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}