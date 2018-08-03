package whaledto

/**
 * Api Url : message/list (GET) 玩家訊息
 */
type MessageList struct {
}


/**
 * Api Url : message/read (POST) 已讀訊息
 */
type MessageRead struct {
	Msg_Id int `json:"msg_id"`
}

/**
 * Api Url : message (GET) 訊息列表
 */
type MessageBackEndList struct {
	Site_Code string `json:"site_code"`
	Count string `json:"count"`
	Page string `json:"page"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	User string `json:"user"`
}

/**
 * Api Url : message (POST) 新增訊息
 */
type CreateMessage struct {
	Site_Code string `json:"site_code"`
	Title string `json:"title"`
	Content string `json:"content"`
	Type int `json:"type"`
	Level string `json:"level"`
	User string `json:"user"`
}

/**
 * Api Url : message/data (GET) 單筆訊息
 */
type MessageData struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
}

/**
 * Api Url : message/user_data (GET) 單筆訊息-玩家資料
 */
type MessageUserData struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
	Count string `json:"count"`
	Page string `json:"page"`
}