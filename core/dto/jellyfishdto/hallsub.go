package jellyfishdto

/**
 * Api Url : hallSub (POST)
 */
type CreateHallSub struct {
	Hall_Code string `json:"hall_code"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Status int `json:"status"`
	Authority int `json:"authority"`
}

/**
 * Api Url : hallSub (PUT)
 */
type UpdataHallSub struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Name string `json:"name"`
	Authority int `json:"authority"`
	Status int `json:"status"`
}

/**
 * Api Url : hallSub (GET)
 */
type HallSubList struct {
	Hall_Code string `json:"hall_code"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : hallSub/data (GET)
 */
type HallSubData struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * Api Url : hallSub/password (GET) 更新廳主子帳號密碼
 */
type UpdataHallSubPassword struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Password string `json:"password"`
}