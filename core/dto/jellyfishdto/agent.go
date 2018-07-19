package jellyfishdto

/**
 * Api Url : Agent (GET)
 */
type AgentList struct {
	//Hall_Code string `json:"hall_code"`
	GeneralAgent_Id string `json:"generalAgent_id"`
	Status string `json:"status"`
	Name string `json:"name"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : agent/data (GET)
 */
type AgentData struct {
	Id string `json:"id"`
}

/**
 * Api Url : agent (POST)
 */
type CreateAgent struct {
	//Hall_Code string `json:"hall_code"`
	GeneralAgent_Id int `json:"generalAgent_id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Take_Part int `json:"take_part"`
}

/**
 * Api Url : agent (PUT)
 */
type UpdateAgent struct {
	//Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Name string `json:"name"`
	Take_Part int `json:"take_part"`
}

/**
 * Api Url : agent/password 更新代理密碼
 */
type UpdateAgentPassword struct {
	//Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Password string `json:"password"`
}

/**
 * Api Url : agent/dropdownlist 取得代理清單(下拉選單用)
 */
type AgentDropdownList struct {
	//Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
}