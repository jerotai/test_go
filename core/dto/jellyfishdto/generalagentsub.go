package jellyfishdto

/**
 * Api Url : generalAgentSub (GET)
 */
type GeneralAgentSubList struct {
	Hall_Code string `json:"hall_code"`
	Status string `json:"status"`
	Account string `json:"account"`
	Name string `json:"name"`
	Page string `json:"page"`
	Count string `json:"count"`
}


/**
 * Api Url : generalAgentSub/data (GET)
 */
type GeneralAgentSubData struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * Api Url : generalAgentSub (POST)
 */
type CreateGeneralAgentSub struct {
	Hall_Code string `json:"hall_code"`
	GeneralAgent_Id int `json:"generalAgent_id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Authority int `json:"authority"`
}

/**
 * Api Url : generalAgentSub (PUT)
 */
type UpdataGeneralAgentSub struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Name string `json:"name"`
	Authority int `json:"authority"`
	Status int `json:"status"`
}

/**
 * Api Url : GeneralAgentSub/password 更新總代子帳號密碼
 */
type UpdataGeneralAgentSubPassword struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Password string `json:"password"`
}