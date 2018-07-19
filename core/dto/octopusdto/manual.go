package octopusdto

/**
 * Api Url : manual
 */
type ManualDepositsList struct {
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
	Txnid string `json:"txnid"`
}

/**
 * Api Url : manualDeposits/pass 人工通過
 */
type ManualDepositsPass struct {
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
	Status string `json:"status"`
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
	Trans_Id string `json:"trans_id"`
	User_Bank_Id string `json:"user_bank_id"`
	User string `json:"user"`
}

/**
 * Api Url : manualDeposits/unlock (PUT) 人工入款列表頁解鎖
 */
type ManualDepositsUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}


/**
 * Api Url : manualWithdraws/unlock (PUT) 人工出款列表頁解鎖
 */
type ManualWithdrawsUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}

/**
 * Api Url : manualWithdraws/audit (GET)
 */
type ManualWithdrawsAudit struct {
	Txnid string `json:"txnid"`
	Type string `json:"type"`
}

/**
 * Api Url : manualWithdraws/audit/pass (POST) 人工出款稽核通過
 */
type ManualWithdrawsAuditPass struct {
	Txnid string `json:"txnid"`
	Cost int `json:"cost"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : manualWithdraws/audit/reject (POST) 人工出款稽核拒絕
 */
type ManualWithdrawsAuditReject struct {
	Txnid string `json:"txnid"`
	Cost int `json:"cost"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : manualWithdraws/grant (GET) 取得人工出款核發對應資訊
 */
type ManualWithdrawsGrant struct {
	Txnid string `json:"txnid"`
	Type string `json:"type"`
}

/**
 * Api Url : manualWithdraws/grant/pass (POST) 人工出款核發通過
 */
type ManualWithdrawsGrantPass struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : manualWithdraws/grant/reject (POST) 人工出款核發拒絕
 */
type ManualWithdrawsGrantReject struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}