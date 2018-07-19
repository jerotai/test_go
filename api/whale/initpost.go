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
		apiDto = &dto.UpdateSite{}
		apiRequestUrl = "site/update"
	case "/banner":
		apiDto = &dto.CreateBanner{}
		apiRequestUrl = "banner"
	case "/banner/update":
		apiDto = &dto.UpdateBanner{}
		apiRequestUrl = "banner/update"
	case "/promotion":
		apiDto = &dto.CreatePromotion{}
		apiRequestUrl = "promotion"
	case "/promotion/update":
		apiDto = &dto.UpdatePromotion{}
		apiRequestUrl = "promotion/update"
	case "/article":
		apiDto = &dto.CreateArticle{}
		apiRequestUrl = "article"
	case "/image":
		apiDto = &dto.CreateImage{}
		apiRequestUrl = "image"
	case "/image/update":
		apiDto = &dto.UpdateImage{}
		apiRequestUrl = "image/update"
	case "/announcement":
		apiDto = &dto.CreateAnnouncement{}
		apiRequestUrl = "announcement"
	case "/bookmark/type":
		apiDto = &dto.CreateBookmarkType{}
		apiRequestUrl = "bookmark/type"
	case "/bookmark/focus":
		apiDto = &dto.CreateBookmarkFocus{}
		apiRequestUrl = "bookmark/focus"
	case "/bookmark":
		apiDto = &dto.CreateBookmark{}
		apiRequestUrl = "bookmark"
	case "/bulletin":
		apiDto = &dto.CreateBulletin{}
		apiRequestUrl = "bulletin"
	case "/page_article":
		apiDto = &dto.CreatePageArticle{}
		apiRequestUrl = "page_article"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}