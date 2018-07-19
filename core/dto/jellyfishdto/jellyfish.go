package jellyfishdto

type LoginRes struct {
	Code string `json:"code"`
	Result struct{
		Api_Token string `json:"api_token"`
		Role_Id int `json:"role_id"`
	} `json:"result"`
	Message string `json:"message"`
}

/**
 * Api Url : login
 */
type Login struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : password 更新自己密碼
 */
type UpdatePassWord struct {
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}

/**
 * Api Url : role 取得角色清單
 */
type RoleList struct {
}

/**
 * api Url : password
 */
type UpdateRolePassword struct {
	Role_Id int `json:"role_id"`
	Id int `json:"id"`
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}