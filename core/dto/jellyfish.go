package dto

/**
 * Api Url : login
 */
type Login struct {
	Account string `json:"account"`
	Password string `json:"password"`
	HallCode string `json:"hall_code"`
}

/**
 * Api Url : hall
 */
type Hall struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	NewHallCode string `json:"new_hall_code"`
	HallCode string `json:"hall_code"`
}

/**
 * Api Url : UpdataHall
 */
type UpdataHall struct {
	Id string `json:"id"`
	Name string `json:"name"`
	HallCode string `json:"hall_code"`
}

/**
 * Api Url : hall/enabled
 */
type EnabledHall struct {
	Id string `json:"id"`
	HallCode string `json:"hall_code"`
}

/**
 * Api Url : hall/subList
 */
type SubList struct {
	Page string `json:"page"`
	Count string `json:"count"`
	HallCode string `json:"hall_code"`
}