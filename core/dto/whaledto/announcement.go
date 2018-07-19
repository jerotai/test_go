package whaledto

/**
 * Api Url : announcement (GET) 系統公告列表
 */
type AnnouncementList struct {
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : announcement/data (GET) 系統公告單筆資料
 */
type AnnouncementData struct {
	Id string `json:"id"`
}

/**
 * Api Url : announcement/hall (GET) 下層廳主列表
 */
type AnnouncementHall struct {
}

/**
 * Api Url : announcement (POST) 新增系統公告
 */
type CreateAnnouncement struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Published_For string `json:"published_for"`
	Published_At string `json:"published_at"`
	Status int `json:"status"`
}

/**
 * Api Url : announcement (PUT) 更新系統公告
 */
type UpdateAnnouncement struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Published_For string `json:"published_for"`
	Published_At string `json:"published_at"`
	Status int `json:"status"`
}

/**
 * Api Url : announcement (DELETE) 刪除系統公告
 */
type DeleteAnnouncement struct {
	Id int `json:"id"`
}

/**
 * Api Url : announcement/batch (DELETE) 批量刪除系統公告
 */
type DeleteAnnouncementBatch struct {
	Ids string `json:"id"`
}
