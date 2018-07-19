package whaledto

import "net/http"

/**
 * Api Url : image (GET) 取得圖庫清單
 */
type ImageList struct {
}

/**
 * Api Url : image (POST) 新增圖庫
 */
type CreateImage struct {
	Name string `json:"name"`
	Image http.File `json:"image"`
}

/**
 * Api Url : image (POST)
 */
type UpdateImage struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Image http.File `json:"image"`
}

/**
 * Api Url : image (DELETE) 刪除圖庫
 */
type DeleteImage struct {
	Id int `json:"id"`
}