package whale

import (
	dto "Stingray/core/dto/whaledto"
	"Stingray/helper"
)

/**
 * init Api Post Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Whale) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/site":
		apiDto = &dto.CreateSite{}
		apiRequestUrl = "site"
	case "/site/update":
		apiDto = &dto.UpdataSite{}
		apiRequestUrl = "site/update"
	case "/banner":
		apiDto = &dto.CreateBanner{}
		apiRequestUrl = "banner"
	case "/banner/update":
		apiDto = &dto.UpdataBanner{}
		apiRequestUrl = "banner/update"
	case "/promotion":
		apiDto = &dto.CreatePromotion{}
		apiRequestUrl = "promotion"
	case "/promotion/update":
		apiDto = &dto.UpdataPromotion{}
		apiRequestUrl = "promotion/update"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}