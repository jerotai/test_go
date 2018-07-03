package jellyfishdto


/**
 * Api Url : hall (GET)
 */
type CreateHall struct {
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

/**
 * Api Url : hall/password 更新下層廳主密碼
 */
type UpdataHallPassword struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Password string `json:"password"`
}

/**
 * Api Url : hall/dropdownlist 取得下層廳主清單(下拉選單用)
 */
type HallDropdownList struct {
	Hall_Code string `json:"hall_code"`
}