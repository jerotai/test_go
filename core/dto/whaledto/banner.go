package whaledto

import (
	"net/http"
)

/**
 * Service Api Url : banner/list (GET) 取得前台Banner清單
 */
type BannerList struct {
}

/**
 * Service Api Url : banner (GET) 取得後台Banner清單
 */
type BackEndBannerList struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : banner/data (GET) 取得後台Banner單筆資料
 */
type BannerData struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
}

/**
 * Api Url : banner (POST) 新增Banner
 */
type CreateBanner struct {
	Site_Code string `json:"site_code"`
	Image http.File `json:"image"`
	Href string `json:"href"`
	Target_Blank int `json:"target_blank"`
	Publish int `json:"publish"`
	Order int `json:"order"`
	Published_At string `json:"published_at"`
}

/**
 * Api Url : banner (PUT) 更新Banner
 */
type UpdateBanner struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Image http.File `json:"image"`
	Href string `json:"href"`
	Target_Blank int `json:"target_blank"`
	Publish int `json:"publish"`
	Order int `json:"order"`
	Published_At string `json:"published_at"`
}

/**
 * Api Url : banner (DELETE) 刪除Banner
 */
type DeleteBanner struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}