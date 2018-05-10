package dto

/**
 * Api Url : login
 */
type Login struct {
	Code string `json:"account"`
	Result string `json:"password"`
	HallCode string `json:"hall_code"`
}

/**
 * Api Url : login Response
 */
type LoginResponse struct {
	Code string `json:"code"`
	Result string `json:"result"`
}
