package whaledto

import "net/http"

/**
 * Service Api Url : promotion/list (GET) 取得前台優惠活動清單
 */
type PromotionList struct {
}

/**
 * Service Api Url : promotion/info (GET) 取得前台優惠活動單筆資料
 */
type PromotionInfo struct {
	Id string `json:"id"`
}

/**
 * Api Url : promotion (GET) 取得優惠活動清單
 */
type PromotionBackEndList struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : promotion/data (GET) 取得優惠活動單筆資料
 */
type PromotionData struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
}

/**
 * Api Url : promotion (POST) 新增優惠活動
 */
type CreatePromotion struct {
	Site_Code string `json:"site_code"`
	Image_Desktop http.File `json:"image_desktop"`
	Image_Mobile http.File `json:"image_mobile"`
	Img_Order int `json:"img_order"`
	Publish int `json:"publish"`
	Published_At string `json:"published_at"`
}

/**
 * Api Url : promotion (PUT) 更新優惠活動
 */
type UpdatePromotion struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Image_Desktop http.File `json:"image_desktop"`
	Image_Mobile http.File `json:"image_mobile"`
	Img_Order int `json:"img_order"`
	Publish int `json:"publish"`
	Published_At string `json:"published_at"`
}

/**
 * Api Url : promotion (DELETE) 刪除優惠活動
 */
type DeletePromotion struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}