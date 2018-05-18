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
	Status int `json::"status"`
}

/**
 * Api Url : hallSub (GET)
 */
type HallSubList struct {
	Hall_Code string `json:"hall_code"`
	Page int `json:"page"`
	Count int `json:"count"`
}