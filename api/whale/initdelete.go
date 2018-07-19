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
	case "/article":
		apiDto = &dto.DeleteArticle{}
		apiRequestUrl = "article"
	case "/image":
		apiDto = &dto.DeleteImage{}
		apiRequestUrl = "image"
	case "/announcement":
		apiDto = &dto.DeleteAnnouncement{}
		apiRequestUrl = "announcement"
	case "/announcement/batch":
		apiDto = &dto.DeleteAnnouncementBatch{}
		apiRequestUrl = "announcement/batch"
	case "/bookmark/focus":
		apiDto = &dto.DeleteBookmarkFocus{}
		apiRequestUrl = "bookmark/focus"
	case "/bookmark/type":
		apiDto = &dto.DeleteBookmarkType{}
		apiRequestUrl = "bookmark/type"
	case "/bookmark":
		apiDto = &dto.DeleteBookmark{}
		apiRequestUrl = "bookmark"
	case "/bulletin":
		apiDto = &dto.DeleteBulletin{}
		apiRequestUrl = "bulletin"
	case "/bulletin/batch":
		apiDto = &dto.BulletinBatch{}
		apiRequestUrl = "bulletin/batch"
	case "/page_article":
		apiDto = &dto.DeletePageArticle{}
		apiRequestUrl = "page_article"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whale InitDeleteApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}