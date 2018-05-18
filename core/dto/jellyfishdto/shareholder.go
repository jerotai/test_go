package jellyfishdto

/**
 * Api Url : shareholder
 */
type ShareholderList struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Status string `json:"status"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Name string `json:"name"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : shareholder/data
 */
type ShareholderData struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
}

/**
 * Api Url : shareholder (POST)
 */
type CreateShareholder struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Take_Part string `json:"take_part"`
}

/**
 * Api Url : shareholder (PUT)
 */
type UpdataShareholder struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Name string `json:"name"`
	Take_Part string `json:"take_part"`
}
