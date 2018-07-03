package octopusdto

/**
 * Api Url : manual
 */
type ManualDepositsList struct {
	Hall_Code string `json:"hall_code"`
	Status string `json:"status"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	Trans_Id string `json:"trans_id"`
	User string `json:"user"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : manual (POST)
 */
type CreateManualDeposits struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Account string `json:"account"`
	Balance int `json:"balance"`
	Bonus int `json:"bonus"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : manualDeposits/audit 人工稽核
 */
type ManualDepositsAudit struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
}

/**
 * Api Url : manualDeposits/pass 人工通過
 */
type ManualDepositsPass struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
	Dama int `json:"dama"`
	Bonus_Dama int `json:"bonus_dama"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : manualDeposits/reject 人工拒絕
 */
type ManualDepositsReject struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
	Dama int `json:"dama"`
	Bonus_Dama int `json:"bonus_dama"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}


/**
 * Api Url : manualWithdraws 人工出款 (POST)
 */
type CreateManualWithdraws struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Account string `json:"account"`
	User_Bank_Id int `json:"user_bank_id"`
	Balance int `json:"balance"`
	Fee float32 `json:"fee"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : manualWithdraws 人工出款列表 (GET)
 */
type ManualWithdrawsList struct {
	Hall_Code string `json:"hall_code"`
	Status int `json:"status"`
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
	Trans_Id string `json:"trans_id"`
	User_Bank_Id int `json:"user_bank_id"`
	User string `json:"user"`
}

/**
 * Api Url : manualDeposits/unlock (PUT) 人工入款列表頁解鎖
 */
type ManualDepositsUnlock struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
	Type int `json:type`
}


/**
 * Api Url : manualWithdraws/unlock (PUT) 人工出款列表頁解鎖
 */
type ManualWithdrawsUnlock struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
	Type int `json:type`
}