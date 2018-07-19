package whitebaitdto

/**
 * Api Url : user_level (GET)
 */
type UserLevel struct {
	Site_Code string `json:"site_code"`
	Name string `json:"name"`
}

/**
 * Api Url : user_level/data (GET)
 */
type UserLevelData struct {
	Id string `json:"id"`
}

/**
 * Api Url : user_level (POST)
 */
type CreateUserLevel struct {
	Site_Code string `json:"site_code"`
	Name string `json:"name"`
	Deposit_Times int `json:"deposit_times"`
	Deposit_Total int `json:"deposit_total"`
	Deposit_Maximum int `json:"deposit_maximum"`
	Withdraw_Times int `json:"withdraw_times"`
	Withdraw_Total int `json:"withdraw_total"`
	Remark string `json:"remark"`
}

/**
 * Api Url : user_level (PUT)
 */
type UpdateUserLevel struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Deposit_Times int `json:"deposit_times"`
	Deposit_Total int `json:"deposit_total"`
	Deposit_Maximum int `json:"deposit_maximum"`
	Withdraw_Times int `json:"withdraw_times"`
	Withdraw_Total int `json:"withdraw_total"`
	Remark string `json:"remark"`
}

/**
 * Api Url : user_level (DELETE)
 */
type DeleteUserLevel struct {
	Id int `json:"id"`
}

/**
 * Api Url : user_level/amount (GET)
 */
type UserLevelAmount struct {
	Id string `json:"id"`
	Type string `json:"type"`
}

/**
 * user_level/amount (PUT)
 */
type UpdateUserLevelAmount struct {
	Id int `json:"id"`
	Deposit_Maximum int `json:"deposit_maximum"`
	Deposit_Minimum int `json:"deposit_minimum"`
	Withdraw_Maximum int `json:"withdraw_maximum"`
	Withdraw_Minimum int `json:"withdraw_minimum"`
	Bonus_Threshold int `json:"bonus_threshold"`
	Dama int `json:"dama"`
	Bonus_Dama int `json:"bonus_dama"`
	Bonus_Ratio int `json:"bonus_ratio"`
	Bonus_Max int `json:"bonus_max"`
	Cost_Ratio int `json:"cost_ratio"`
	Fee_Ratio int `json:"fee_ratio"`
	Fee_Maximum int `json:"fee_maximum"`
	Fee_Time_range int `json:"fee_time_range"`
	Fee_Times int `json:"fee_times"`
}

/**
 * user_level/payment
 */
type UserLevelPayment struct {
	Id string `json:"id"`
}

/**
 * user_level/payment (PUT)
 */
type UpdateUserLevelPayment struct {
	Id int `json:"id"`
	Fourth_Ids string `json:"fourth_ids"`
}

/**
 * Api Url : user_level/company_bank (GET)
 */
type UserLevelCompanyBankList struct {
	Id string `json:"id"`
}

/**
 * user_level/company_bank (PUT)
 */
type UpdateUserLevelCompanyBank struct {
	Id int `json:"id"`
	Company_Bank_Ids string `json:"company_bank_ids"`
}

/**
 * Api Url : user_level/user (PUT) 異動會員的會員層級
 */
type UpdateBackEndUserLevel struct {
	User_Ids string `json:"user_ids"`
	User_Level_Id int `json:"user_level_id"`
}

/**
 * Api Url : user_level/batch (PUT) 異動會員層級的會員
 */
type UpdateBackEndUserLevelBatch struct {
	Before_User_Level_Id int `json:"before_user_level_id"`
	After_User_Level_Id int `json:"after_user_level_id"`
}

/**
 * Api Url : user_level/user_list (GET) 根據會員層級取得會員清單資料
 */
type BackEndUserLevelUserList struct {
	Id string `json:"id"`
	Account string `json:"account"`
	Lock_Rank string `json:"lock_rank"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : user_level/dropdownlist (GET) 取得會員層級清單（下拉選單）
 */
type UserLevelDropdownList struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : user_level/user/preview (GET) 異動會員的會員層級，預覽
 */
type UserLevelUserPreview struct {
	User_Ids string `json:"user_ids"`
	User_Level_Id string `json:"user_level_id"`
}

/**
 * Api Url : user_level/batch/preview (GET) 異動會員層級的會員，預覽
 */
type UserLevelBatchPreview struct {
	Before_User_Level_Id string `json:"before_user_level_id"`
	After_User_Level_Id string `json:"after_user_level_id"`
	
}