package lobsterdto

/**
 * Api Url : userLoginReport/user_login_drive_count (GET) 登入裝置報表
 */
type UserLoginDriveCount struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}

/**
 * Api Url : userLoginReport/user_login_info (GET) 會員登入資訊
 */
type UserLoginInfo struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
	Ip_Is_Repeat string `json:"ip_is_repeat"`
	Account string `json:"account"`
	Ip string `json:"ip"`
}

type UserLoginInfoOnline struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
	Ip_Is_Repeat string `json:"ip_is_repeat"`
	Account string `json:"account"`
	Ip string `json:"ip"`
}

/**
 * Api Url : userLoginReport/user_login_record (GET) 會員登入紀錄
 */
type UserLoginRecord struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
	Account string `json:"account"`
	Ip string `json:"ip"`
}