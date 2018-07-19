package whaledto

/**
 * Api Url : bookmark/list (GET) 取得頁籤清單
 */
type BookmarkList struct {
}

/**
 * Api Url : bookmark/focus (POST) 新增關注頁籤的功能
 */
type CreateBookmarkFocus struct {
	Action_Id int `json:"action_id"`
	Content interface {} `json:"content"`
	Image_Name string `json:"image_name"`
	Name string `json:"name"`
}

/**
 * Api Url : bookmark/focus (DELETE) 刪除關注頁籤的功能
 */
type DeleteBookmarkFocus struct {
	Id int `json:"id"`
}

/**
 * Api Url : bookmark/type (GET) 取得站台頁籤類型
 */
type BookmarkType struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : bookmark/type (POST) 新增站台頁籤類型
 */
type CreateBookmarkType struct {
	Site_Code string `json:"site_code"`
	Page_Name string `json:"page_name"`
}

/**
 * Api Url : bookmark/type (PUT) 更新站台頁籤類型
 */
type UpdateBookmarkType struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Page_Name string `json:"page_name"`
}

/**
 * Api Url : bookmark/type (DELETE) 刪除站台頁籤類型
 */
type DeleteBookmarkType struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}

/**
 * Api Url : bookmark (GET) 取得站台頁籤清單
 */
type BackEndBookmarkList struct {
	Site_Code string `json:"site_code"`
	Type_Id string `json:"type_id"`
}

/**
 * Api Url : bookmark/data (GET) 取得站台頁籤功能設定
 */
type BookmarkData struct {
	Site_Code string `json:"site_code"`
	Type_Id string `json:"type_id"`
	Id string `json:"id"`
}

/**
 * Api Url : bookmark (POST) 新增站台頁籤功能設定
 */
type CreateBookmark struct {
	Site_Code string `json:"site_code"`
	Type_Id string `json:"type_id"`
	Action_Id int `json:"action_id"`
	Name string `json:"name"`
	Content interface {} `json:"content"`
	Image_Id int `json:"image_id"`
}

/**
 * Api Url : bookmark (PUT) 更新站台頁籤功能設定
 */
type UpdateBookmark struct {
	Site_Code string `json:"site_code"`
	Type_Id string `json:"type_id"`
	Id int `json:"id"`
	Action_Id int `json:"action_id"`
	Content interface {} `json:"content"`
	Image_Id int `json:"image_id"`
	Name string `json:"name"`
}

/**
 * Api Url : bookmark (DELETE) 刪除站台頁籤功能設定
 */
type DeleteBookmark struct {
	Site_Code string `json:"site_code"`
	Type_Id string `json:"type_id"`
	Id int `json:"id"`
}

/**
 * Api Url : bookmark/code (GET) 取得站台頁籤前台代碼
 */
type BookmarkCode struct {
}