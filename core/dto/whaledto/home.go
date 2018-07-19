package whaledto

/**
 * Api Url : home (GET) 首頁
 */
type Home struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}

/**
 * Api Url : home\announcement (GET) 首頁系統公告列表
 */
type HomeAnnouncement struct {
	Count string `json:"count"`
	Page string `json:"page"`
}

/**
 * Api Url : home\announcement\data (GET) 首頁系統公告-單筆
 */
type HomeAnnouncementData struct {
	Id string `json:"id"`
}