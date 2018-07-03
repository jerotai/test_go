package jellyfishdto

/**
 * Api Url : shareholderSub (GET)
 */
type ShareholderSubList struct {
	Hall_Code string `json:"hall_code"`
	Page string `json:"page"`
	Count string `json:"count"`
	Status string `json:"status"`
	Account string `json:"account"`
	Name string `json:"name"`
}

/**
 * Api Url : shareholderSub/data (GET)
 */
type ShareholderSubData struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * Api Url : shareholderSub (POST)
 */
type CreateShareholderSub struct {
	Hall_Code string `json:"hall_code"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Authority int `json:"authority"`
}

/**
 * Api Url : shareholderSub (PUT)
 */
type UpdataShareholderSub struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Name string `json:"name"`
	Authority int `json:"authority"`
	Status int `json:"status"`
}

/**
 * Api Url : shareholderSub/password (PUT) 更新股東密碼
 */
type UpdataShareholderSubPassword struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Password string `json:"password"`
}
