package jellyfishdto

/**
 * Api Url : shareholder (GET)
 */
type ShareholderList struct {
	Site_Code string `json:"site_code"`
	Status string `json:"status"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Name string `json:"name"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : shareholder/data (GET)
 */
type ShareholderData struct {
	Id string `json:"id"`
}

/**
 * Api Url : shareholder (POST)
 */
type CreateShareholder struct {
	Site_Code string `json:"site_code"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Take_Part int `json:"take_part"`
}

/**
 * Api Url : shareholder (PUT)
 */
type UpdateShareholder struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Take_Part int `json:"take_part"`
}

/**
 * Api Url : shareholder/password (PUT) 更新股東密碼
 */
type UpdateShareholderPassword struct {
	Id int `json:"id"`
	Password string `json:"password"`
}

/**
 * Api Url : shareholder/dropdownlist 取得股東清單(下拉選單用)
 */
type ShareholderDropdownList struct {
	Site_Code string `json:"site_code"`
}