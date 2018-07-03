package jellyfishdto

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
type UpdataPassWord struct {
	Hall_Code string `json:"hall_code"`
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}

/**
 * Api Url : role 取得角色清單
 */
type RoleList struct {
	Hall_Code string `json:"hall_code"`
}

/**
 * api Url : password
 */
type UpdataRolePassword struct {
	Hall_Code string `json:"hall_code"`
	Role_Id int `json:"role_id"`
	Id int `json:"id"`
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}