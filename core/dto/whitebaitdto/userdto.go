package whitebaitdto

/**
 * Api Url : user/login (POST) 登入
 */
type UserLoginRes struct {
	Code string `json:"code"`
	Result struct{
		Api_Token string `json:"api_token"`
	} `json:"result"`
	Message string `json:"message"`
}

/**
 * user/register (POST)
 */
type CreateUser struct {
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
}

/**
 * user list (GET)
 */
type UserList struct {
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
	Email string `json:"email"`
	Qq string `json:"qq"`
	Wechat string `json:"wechat"`
	User_Bank_Account string `json:"user_bank_account"`
}

/**
 * user login (POST)
 */
type UserLogin struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Plat_Id int `json:"plat_id"`
	Login_Ip string `json:"login_ip"`
	Domain string `json:"domain"`
}

/**
 * user/promo_code
 */
type UserPromoCode struct {
	Promotion_Code string `json:"promotion_code"`
	Promotion_Type string `json:"promotion_type"`
}

/**
 * user/plat
 */
type UserPlat struct {
}

/**
 * user/data
 */
type UserData struct {
}

/**
 * user/balance
 */
type UserBalance struct {
}

/**
 * user/password
 */
type UpdateUserPassword struct {
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}

/**
 * user/company_bank
 */
type UserCompanyBank struct {
}

/**
 * user/info (GET)
 */
type UserInfo struct {
	Id string `json:"id"`
}

/**
 * user/info (PUT) 更新會員單筆資料
 */
type UpdateUserInfo struct {
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
	Site_Code string `json:"site_code"`
}

/**
 * user/register/setting (PUT)
 */
type UpdateUserRegisterSetting struct {
	Site_Code string `json:"site_code"`
	Code string `json:"code"`
}

/**
 * Api Url : user/alive 取得連線狀態
 */
type UserAlive struct {
}

/**
 * Api Url : user/password/withdraw/check (GET) 確認是否設定取款密碼
 */
type UserPasswordWithdrawCheck struct {
}

/**
 * Api Url : user/password/withdraw (POST) 新增取款密碼
 */
type CreateUserPasswordWithdraw struct {
	Password string `json:"password"`
}

/**
 * Api Url : user/password/withdraw (PUT) 新增取款密碼
 */
type UpdateUserPasswordWithdraw struct {
	Password string `json:"password"`
	New_Password string `json:"new_password"`
}

/**
 * Api Url : user/password/withdraw/update (PUT) 新增取款密碼
 */
type UpdateBackEndUserPasswordWithdraw struct {
	Password string `json:"password"`
	User_Id int `json:"user_id"`
}

/**
 * Api Url : user/create (POST) 後台 - 新增會員
 */
type BackEndCreateUser struct {
	Site_Code string `json:"site_code"`
	Ag_Id int `json:"ag_id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Lock_Rank int `json:"lock_rank"`
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
type UpdateBackEndUserPassword struct {
	Id int `json:"id"`
	Password string `json:"password"`
}

/**
 * Api Url : user/blurry (GET) 模糊搜尋會員帳號
 */
type UserBlurry struct {
	Site_Code string `json:"site_code"`
	Account string `json:"account"`
}