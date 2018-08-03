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
	case "/article/list":
		apiDto = &dto.ArticleList{}
		apiRequestUrl = "article"
	case "/article/page":
		apiDto = &dto.ArticlePage{}
		apiRequestUrl = "article/page"
	case "/article/data":
		apiDto = &dto.ArticleData{}
		apiRequestUrl = "article/data"
	case "/article/info":
		apiDto = &dto.ArticleInfo{}
		apiRequestUrl = "article/info"
	case "/image/list":
		apiDto = &dto.ImageList{}
		apiRequestUrl = "image"
	case "/image/page":
		apiDto = &dto.ImagePage{}
		apiRequestUrl = "image/page"
	case "/home/list":
		apiDto = &dto.Home{}
		apiRequestUrl = "home"
	case "/home/announcement":
		apiDto = &dto.HomeAnnouncement{}
		apiRequestUrl = "home/announcement"
	case "/home/announcementData":
		apiDto = &dto.HomeAnnouncementData{}
		apiRequestUrl = "home/announcement/data"
	case "/announcement/list":
		apiDto = &dto.AnnouncementList{}
		apiRequestUrl = "announcement"
	case "/announcement/data":
		apiDto = &dto.AnnouncementData{}
		apiRequestUrl = "announcement/data"
	case "/announcement/hall":
		apiDto = &dto.AnnouncementHall{}
		apiRequestUrl = "announcement/hall"
	case "/bookmark/list":
		apiDto = &dto.BookmarkList{}
		apiRequestUrl = "bookmark/list"
	case "/bookmark/type":
		apiDto = &dto.BookmarkType{}
		apiRequestUrl = "bookmark/type"
	case "/bookmark/backEndList":
		apiDto = &dto.BackEndBookmarkList{}
		apiRequestUrl = "bookmark"
	case "/bookmark/data":
		apiDto = &dto.BookmarkData{}
		apiRequestUrl = "bookmark/data"
	case "/bookmark/code":
		apiDto = &dto.BookmarkCode{}
		apiRequestUrl = "bookmark/code"
	case "/bulletin/list":
		apiDto = &dto.BulletinList{}
		apiRequestUrl = "bulletin/list"
	case "/bulletin/backend_list":
		apiDto = &dto.BackEndBulletinList{}
		apiRequestUrl = "bulletin"
	case "/bulletin/data":
		apiDto = &dto.BulletinData{}
		apiRequestUrl = "bulletin/data"
	case "/page_article/site":
		apiDto = &dto.PageArticle{}
		apiRequestUrl = "page_article"
	case "/page_article/info":
		apiDto = &dto.PageArticleInfo{}
		apiRequestUrl = "page_article/info"
	case "/page_article/code":
		apiDto = &dto.PageArticleCode{}
		apiRequestUrl = "page_article/code"
	case "/message/list":
		apiDto = &dto.MessageList{}
		apiRequestUrl = "message/list"
	case "/message/backEndList":
		apiDto = &dto.MessageBackEndList{}
		apiRequestUrl = "message"
	case "/message/data":
		apiDto = &dto.MessageData{}
		apiRequestUrl = "message/data"
	case "/message/userData":
		apiDto = &dto.MessageUserData{}
		apiRequestUrl = "message/user_data"
	case "/customer_service/url":
		apiDto = &dto.CustomerServiceUrl{}
		apiRequestUrl = "customer_service"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whale InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
