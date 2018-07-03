package jellyfishdto

/**
 * Api Url : authGroup (POST)
 */
type CreateAuthGroup struct {
	Hall_Code string `json:"hall_code"`
	Auth_Json string `json:"auth_json"`
	Role_Id int `json:"role_id"`
	Name string `json:"name"`
}

/**
 * Api Url : authGroup (PUT)
 */
type UpdataAuthGroup struct {
	Hall_Code string `json:"hall_code"`
	Auth_Json string `json:"auth_json"`
	Group_Id int `json:"group_id"`
	Name string `json:"name"`
}

/**
 * Api Url : authGroup (GET)
 */
type AuthGroupList struct {
	Hall_Code string `json:"hall_code"`
	Count string `json:"count"`
	Page string `json:"page"`
}

/**
 * Api Url : authGroup/detail (GET)
 */
type AuthGroupDatail struct {
	Hall_Code string `json:"hall_code"`
	Group_Id string `json:"group_id"`
}

/**
 * Api Url : authGroup/detail (GET)
 */
type AuthGroupDropDown struct {
	Hall_Code string `json:"hall_code"`
}