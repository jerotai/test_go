package whaledto

import "net/http"

/**
 * API Url site
 */
type SiteList struct {
	Hall_Code string `json:"hall_code"`
	Search_Hall_Code string `json:"search_hall_code"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * API Url site/data
 */
type SiteData struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * API Url site (POST)
 */
type CreateSite struct {
	Hall_Code string `json:"hall_code"`
	Name string `json:"name"`
	Title string `json:"title"`
	Logo http.File `json:"logo"`
	Customer_Url string `json:"customer_url"`
	Site_Code string `json:"site_code"`
	Domains string `json:"domains"`
}

/**
 * API Url site (POST) 更新子站資料
 */
type UpdataSite struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Logo http.File `json:"logo"`
	Customer_Url string `json:"customer_url"`
}

/**
 * API Url site/enabled (PUT)
 */
type SiteEnabled struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
}

/**
 * API Url site/agent/default (PUT)
 */
type SiteAgentDefault struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Ag_Id int `json:"ag_id"`
}

/**
 * API Url site/dropdown 取得子站清單(下拉選單用)
 */
type SiteDropdownList struct {
	Hall_Code string `json:"hall_code"`
	Search_Hall_Code string `json:"search_hall_code"`
	Status string `json:"status"`
}