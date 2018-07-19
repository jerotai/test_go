package whaledto

/**
 * Api Url : bulletin/list (GET) 公告清單
 */
type BulletinList struct {
}

/**
 *  Api Url : bulletin (GET) 公告列表
 */
type BackEndBulletinList struct {
	Site_Code string `json:"site_code"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : bulletin/data (GET) 公告單筆資料
 */
type BulletinData struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
}

/**
 * Api Url : bulletin (POST) 新增公告
 */
type CreateBulletin struct {
	Site_Code string `json:"site_code"`
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
	Content string `json:"content"`
	Pop_Out int `json:"pop_out"`
	Marquee int `json:"marquee"`
	Priority int `json:"priority"`
	Publish int `json:"publish"`
	Published_At string `json:"published_at"`
}

/**
 *  Api Url : bulletin (PUT) 更新公告
 */
type UpdateBulletin struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
	Content string `json:"content"`
	Pop_Out int `json:"pop_out"`
	Marquee int `json:"marquee"`
	Priority int `json:"priority"`
	Publish int `json:"publish"`
	Published_At string `json:"published_at"`
}

/**
 * Api Url : bulletin (DELETE) 刪除公告
 */
type DeleteBulletin struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}

/**
 * bulletin/batch (DELETE) 批量刪除公告
 */
type BulletinBatch struct {
	Site_Code string `json:"site_code"`
	Ids string `json:"ids"`
}