package octopusdto

/**
 * Api Url : provider 建立轉帳通知單(轉帳新增)
 */
type CreateProviderDeposits struct {
	Company_Bank_Id int `json:"company_bank_id"`
	Deposit_Time string `json:"deposit_time"`
	User_Name string `json:"user_name"`
	Method string `json:"method"`
	Point string `json:"point"`
	Balance int `json:"balance"`
}

/**
 * Api Url : providerDeposits 轉帳列表
 */
type ProviderDepositsList struct {
	Status string `json:"status"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	Trans_Id string `json:"trans_id"`
	User string `json:"user"`
	Company_Bank_Id string `json:"company_bank_id"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : providerDeposits\audit 轉帳稽核
 */
type ProviderDepositsAudit struct {
	Txnid string `json:"txnid"`
}

/**
 * Api Url : providerDeposits/pass 轉帳通過
 */
type ProviderDepositsPass struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : providerDeposits/reject 轉帳拒絕
 */
type ProviderDepositsReject struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : providerDeposits/limit 取得轉帳入款金額上下限
 */
type ProviderDepositsLimit struct {
}

/**
 * Api Url : providerWithdraws 取得轉帳出款列表 (GET)
 */
type ProviderWithdrawsList struct {
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
 * Api Url : providerWithdraws/limit 取得轉帳出款金額上下限 (GET)
 */
type ProviderWithdrawsLimit struct {
}

/**
 * Api Url : providerWithdraws 建立轉帳出款 (POST)
 */
type CreateProviderWithdraws struct {
	Site_Code string `json:"site_code"`
	User_Bank_Id int `json:"user_bank_id"`
	Password string `json:"password"`
	Balance int `json:"balance"`
}

/**
 * Api Url : providerDeposits/unlock (PUT) 轉帳入款列表頁解鎖
 */
type ProviderDepositsUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}

/**
 * Api Url : providerWithdraws/unlock (PUT) 轉帳出款列表頁解鎖
 */
type ProviderWithdrawsUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}

/**
 * Api Url : providerWithdraws/audit (GET) 轉帳出款稽核
 */
type ProviderWithdrawsAudit struct {
	Txnid string `json:"txnid"`
	Type string `json:"type"`
}

/**
 * Api Url : providerWithdraws/audit/pass (POST) 轉帳出款稽核通過
 */
type ProviderWithdrawsAuditPass struct {
	Txnid string `json:"txnid"`
	Cost int `json:"cost"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : providerWithdraws/audit/reject (POST) 轉帳出款稽核拒絕
 */
type ProviderWithdrawsAuditReject struct {
	Txnid string `json:"txnid"`
	Cost int `json:"cost"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : providerWithdraws/grant (GET) 轉帳出款核發
 */
type ProviderWithdrawsGrant struct {
	Txnid string `json:"txnid"`
	Type string `json:"type"`
}

/**
 * Api Url : providerWithdraws/grant/pass (POST) 轉帳出款核發通過
 */
type ProviderWithdrawsGrantPass struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : providerWithdraws/grant/reject (POST) 轉帳出款核發拒絕
 */
type ProviderWithdrawsGrantReject struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}