package jellyfishdto

/**
 * Api Url : login
 */
type Login struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Hall_Code string `json:"hall_code"`
}
