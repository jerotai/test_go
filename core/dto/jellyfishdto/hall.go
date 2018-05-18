package jellyfishdto


/**
 * Api Url : hall
 */
type Hall struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	NewHallCode string `json:"new_hall_code"`
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : UpdataHall
 */
type UpdataHall struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : hall/enabled
 */
type EnabledHall struct {
	Id int `json:"id"`
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : hall/subList
 */
type SubList struct {
	Page string `json:"page"`
	Count string `json:"count"`
	Hall_Code string `json:"hall_code"`
}
