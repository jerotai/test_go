package whaledto

/**
 * Service Api Url : article (GET) 後台-取得文案清單
 */
type ArticleList struct {
	Site_Code string `json:"site_code"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : article/data (GET) 後台-取得文案單筆資料
 */
type ArticleData struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
}

/**
 * Api Url : `article` (POST) 後台-新增文案
 */
type CreateArticle struct {
	Site_Code string `json:"site_code"`
	Title string `json:"title"`
	Text string `json:"text"`
}

/**
 * Api Url : article (PUT) 後台-編輯文案
 */
type UpdateArticle struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
}

/**
 * Api Url : article (DELETE) 後台-刪除文案
 */
type DeleteArticle struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}

/**
 * Api Url : article/info (GET) 取得前台顯示的文案內容
 */
type ArticleInfo struct {
	Id string `json:"id"`
}

/**
 * Api Url : page_article/info (GET) 取得頁面對應文案
 */
type PageArticleInfo struct {
}

/**
 * Api Url : page_article/code (GET) 取得站台頁面代碼
 */
type PageArticleCode struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : page_article (GET) 取得站台頁面需對應文案的清單
 */
type PageArticle struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : page_article (POST) 新增站台頁面對應文案
 */
type CreatePageArticle struct {
	Site_Code string `json:"site_code"`
	Code string `json:"code"`
	Doc_Id int `json:"doc_id"`
}

/**
 * Api Url : page_article (PUT) 更新站台頁面對應文案
 */
type UpdatePageArticle struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Code string `json:"code"`
	Doc_id int `json:"doc_id"`
}

/**
 * Api Url : page_article (DELETE) 刪除站台頁面對應文案
 */
type DeletePageArticle struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}