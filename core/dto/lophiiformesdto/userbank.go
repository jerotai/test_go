package lophiiformesdto

/**
 * Api Url : user_bank (GET) 取得會員銀行卡清單
 */
type UserBankList struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : user_bank/info (GET) 取得會員銀行卡資料
 */
type UserBankInfo struct {
	Site_Code string `json:"site_code"`
	Id string `json:"id"`
}

/**
 * Api Url : user_bank (POST) 新增會員銀行卡
 */
type CreateUserBank struct {
	Site_Code string `json:"site_code"`
	Bank_Id int `json:"bank_id"`
	Name string `json:"name"`
	Account string `json:"account"`
	Account_Point string `json:"account_point"`
}

/**
 * Api Url : user_bank (PUT) 更新會員銀行卡
 */
type UpdataUserBank struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
	Bank_Id int `json:"bank_id"`
	Name string `json:"name"`
	Account string `json:"account"`
	Account_Point string `json:"account_point"`
}

/**
 * Api Url : user_bank (DELETE) 會員銀行卡
 */
type DeleteUserBank struct {
	Site_Code string `json:"site_code"`
	Id int `json:"id"`
}


/**
 * Api Url : user_bank (GET) 後台取得會員銀行卡清單
 */
type BackEndUserBankList struct {
	Hall_Code string `json:"hall_code"`
	User_Id string `json:"user_id"`
}

/**
 * Api Url : user_bank/info (GET) 後台取得會員銀行卡資料
 */
type BackEndUserBankInfo struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
	User_Id string `json:"user_id"`
}

/**
 * Api Url : user_bank (POST) 後台新增會員銀行卡
 */
type BackEndCreateUserBank struct {
	Hall_Code string `json:"hall_code"`
	User_Id int `json:"user_id"`
	Bank_Id int `json:"bank_id"`
	Name string `json:"name"`
	Account string `json:"account"`
	Account_Point string `json:"account_point"`
}

/**
 * Api Url : user_bank (PUT) 後台更新會員銀行卡
 */
type BackEndUpdataUserBank struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	User_Id int `json:"user_id"`
	Bank_Id int `json:"bank_id"`
	Name string `json:"name"`
	Account string `json:"account"`
	Account_Point string `json:"account_point"`
}

/**
 * Api Url : user_bank (DELETE) 刪除會員銀行卡
 */
type BackEndDeleteUserBank struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	User_Id int `json:"user_id"`
}
