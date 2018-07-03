package whale

import (
	dto "Stingray/core/dto/whaledto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Whale) InitGetApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/menu/list":
		apiDto = &dto.Menu{}
		apiRequestUrl = "menu"
	case "/site/list":
		apiDto = &dto.SiteList{}
		apiRequestUrl = "site"
	case "/site/dropdownlist":
		apiDto = &dto.SiteDropdownList{}
		apiRequestUrl = "site/dropdownlist"
	case "/site/data":
		apiDto = &dto.SiteData{}
		apiRequestUrl = "site/data"
	case "/banner/list":
		apiDto = &dto.BannerList{}
		apiRequestUrl = "banner/list"
	case "/banner/backEndlist":
		apiDto = &dto.BackEndBannerList{}
		apiRequestUrl = "banner"
	case "/banner/data":
		apiDto = &dto.BannerData{}
		apiRequestUrl = "banner/data"
	case "/promotion/list":
		apiDto = &dto.PromotionList{}
		apiRequestUrl = "promotion/list"
	case "/promotion/info":
		apiDto = &dto.PromotionInfo{}
		apiRequestUrl = "promotion/info"
	case "/promotion/backEndList":
		apiDto = &dto.PromotionBackEndList{}
		apiRequestUrl = "promotion"
	case "/promotion/data":
		apiDto = &dto.PromotionData{}
		apiRequestUrl = "promotion/data"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whale InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
