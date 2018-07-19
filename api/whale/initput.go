package whale

import (
	dto "Stingray/core/dto/whaledto"
	"Stingray/helper"
)

/**
 * init Api Put Config
 * dto = &dto.Menu{}
 * apiRequestUrl = "menu"
 *
 */
func (i *Whale) InitPutApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/site/enabled":
		apiDto = &dto.SiteEnabled{}
		apiRequestUrl = "site/enabled"
	case "/site/agent/default":
		apiDto = &dto.SiteAgentDefault{}
		apiRequestUrl = "site/agent/default"
	case "/article":
		apiDto = &dto.UpdateArticle{}
		apiRequestUrl = "article"
	case "/announcement":
		apiDto = &dto.UpdateAnnouncement{}
		apiRequestUrl = "announcement"
	case "/bookmark/type":
		apiDto = &dto.UpdateBookmarkType{}
		apiRequestUrl = "bookmark/type"
	case "/bookmark":
		apiDto = &dto.UpdateBookmark{}
		apiRequestUrl = "bookmark"
	case "/bulletin":
		apiDto = &dto.UpdateBulletin{}
		apiRequestUrl = "bulletin"
	case "/page_article":
		apiDto = &dto.UpdatePageArticle{}
		apiRequestUrl = "page_article"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}