package jellyfishdto

/**
 * Api Url : generalAgent (GET)
 */
type GeneralAgentList struct {
	Shareholder_Id string `json:"shareholder_id"`
	Status string `json:"status"`
	Name string `json:"name"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : generalAgent/data (GET)
 */
type GeneralAgentData struct {
	Id string `json:"id"`
}

/**
 * Api Url : generalAgent (POST)
 */
type CreateGeneralAgent struct {
	Shareholder_Id int `json:"shareholder_id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Take_Part int `json:"take_part"`
}

/**
 * Api Url : generalAgent (PUT)
 */
type UpdateGeneralAgent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Take_Part int `json:"take_part"`
}

/**
 * Api Url : generalAgent/password 更新新總代密碼
 */
type UpdateGeneralAgentPassword struct {
	Id int `json:"id"`
	Password string `json:"password"`
}

/**
 * Api Url : generalAgent/dropdownlist 取得總代理清單(下拉選單用)
 */
type GeneralAgentDropdownList struct {
	Site_Code string `json:"site_code"`
}