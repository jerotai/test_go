package whitebaitdto

/**
 * user/register (POST)
 */
type CreateUser struct {
	Site_Code string `json:"site_code"`
	Promotion_Code string `json:"promotion_code"`
	Promotion_Type string `json:"promotion_type"`
	Account string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Qq string `json:"qq"`
	Wechat string `json:"wechat"`
}

/**
 * user/register (PUT) 取得註冊設定
 */
type UserRegister struct {
	Site_Code string `json:"site_code"`
}

/**
 * user list (GET)
 */
type UserList struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Page string `json:"page"`
	Count string `json:"count"`
	Ag_Id string`json:"ag_id"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Status string `json:"status"`
	Account string `json:"account"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Qq string `json:"qq"`
	Wechat string `json:"wechat"`
	User_Bank_Account string `json:"user_bank_account"`
}

/**
 * user login (POST)
 */
type UserLogin struct {
	Site_Code string `json:"site_code"`
	Account string `json:"account"`
	Password string `json:"password"`
	Plat_Id int `json:"plat_id"`
	Login_Ip string `json:"login_ip"`
}

/**
 * user/promo_code
 */
type UserPromoCode struct {
	Site_Code string `json:"site_code"`
	Promotion_Code string `json:"promotion_code"`
	Promotion_Type string `json:"promotion_type"`
}

/**
 * user/plat
 */
type UserPlat struct {
	Site_Code string `json:"site_code"`
}

/**
 * user/data
 */
type UserData struct {
	Site_Code string `json:"site_code"`
}

/**
 * user/balance
 */
type UserBalance struct {
	Site_Code string `json:"site_code"`
}

/**
 * user/password
 */
type UpdataUserPassword struct {
	Site_Code string `json:"site_code"`
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}

/**
 * user/company_bank
 */
type UserCompanyBank struct {
	Site_Code string `json:"site_code"`
}

/**
 * user/info (GET)
 */
type UserInfo struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * user/info (PUT) 更新會員單筆資料
 */
type UpdataUserInfo struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
	Name string `json:"name"`
	NickName string `json:"nickname"`
	Status int `json:"status"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Qq string `json:"qq"`
	Wechat string `json:"wechat"`
	
}

/**
 * user/register/setting (GET)
 */
type UserRegisterSetting struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
}

/**
 * user/register/setting (PUT)
 */
type UpdataUserRegisterSetting struct {
	Site_Code string `json:"site_code"`
	Code string `json:"code"`
}

/**
 * Api Url : user/alive 取得連線狀態
 */
type UserAlive struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : user/password/withdraw/check (GET) 確認是否設定取款密碼
 */
type UserPasswordWithdrawCheck struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : user/password/withdraw (POST) 新增取款密碼
 */
type CreateUserPasswordWithdraw struct {
	Site_Code string `json:"site_code"`
	Password string `json:"password"`
}

/**
 * Api Url : user/password/withdraw (PUT) 新增取款密碼
 */
type UpdataUserPasswordWithdraw struct {
	Site_Code string `json:"site_code"`
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}

/**
 * Api Url : user/password/withdraw/update (PUT) 新增取款密碼
 */
type UpdataBackEndUserPasswordWithdraw struct {
	Hall_Code string `json:"hall_code"`
	Password string `json:"password"`
	User_Id int `json:"user_id"`
}

/**
 * Api Url : user/create (POST) 後台 - 新增會員
 */
type BackEndCreateUser struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Agent_Id int `json:"agent_id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Qq string `json:"qq"`
	Wechat string `json:"wechat"`
}

/**
 *  Api Url : user/password/update (PUT) 後台-更新會員密碼
 */
type UpdataBackEndUserPassword struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Password string `json:"password"`
}