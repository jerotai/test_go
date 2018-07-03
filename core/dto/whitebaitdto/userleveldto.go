package whitebaitdto

/**
 * Api Url : user_level (GET)
 */
type UserLevel struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Name string `json:"name"`
}

/**
 * Api Url : user_level/data (GET)
 */
type UserLevelData struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * Api Url : user_level (POST)
 */
type CreateUserLevel struct {
	Hall_Code string `json:"hall_code"`
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
type UpdataUserLevel struct {
	Hall_Code string `json:"hall_code"`
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
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
}

/**
 * Api Url : user_level/amount (GET)
 */
type UserLevelAmount struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
	Type string `json:"type"`
}

/**
 * user_level/amount (PUT)
 */
type UpdataUserLevelAmount struct {
	Hall_Code string `json:"hall_code"`
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
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * user_level/payment (PUT)
 */
type UpdataUserLevelPayment struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Fourth_Ids string `json:"fourth_ids"`
}

/**
 * Api Url : user_level/company_bank (GET)
 */
type UserLevelCompanyBankList struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * user_level/company_bank (PUT)
 */
type UpdataUserLevelCompanyBank struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Company_Bank_Ids string `json:"company_bank_ids"`
}

/**
 * Api Url : user_level/user (PUT) 異動會員的會員層級
 */
type UpdataBackEndUserLevel struct {
	Hall_Code string `json:"hall_code"`
	User_Ids string `json:"user_ids"`
	User_Level_Id int `json:"user_level_id"`
}

/**
 * Api Url : user_level/batch (PUT) 異動會員層級的會員
 */
type UpdataBackEndUserLevelBatch struct {
	Hall_Code string `json:"hall_code"`
	Before_User_Level_Id int `json:"before_user_level_id"`
	After_User_Level_Id int `json:"after_user_level_id"`
}