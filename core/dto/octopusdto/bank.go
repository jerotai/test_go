package octopusdto

/**
 * Api Url : bank/list 網銀列表
 */
type BankList struct {
	Page string `json:"page"`
	Count string `json:"count"`
	Status string `json:"status"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	Trans_Id string `json:"trans_id"`
	User string `json:"user"`
	Fourth_Id string `json:"fourth_id"`
	Bank_Id string `json:"bank_id"`
}

/**
 * Api Url : bankDeposits/pass 網銀通過
 */
type BankDepositsPass struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/***
 * Api Url : bankDeposits/menu 網銀下拉式選單
 */
type BankDepositsMenu struct {
	Site_Code string `json:"site_code"`
}

/***
 * Api Url : bankDeposits 網銀新增
 */
type CreateBankDeposits struct {
	Fourth_Id int `json:"fourth_id"`
	Bank_Id int `json:"bank_id"`
	Amount float64 `json:"amount"`
}

/**
 * Api Url : bankDeposits/front_menu
 */
type BankDepositsFrontMenu struct {
}

/**
 * Api Url : bankDeposits/reject 網銀拒絕
 */
type BankDepositsReject struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : bankDeposits/audit 網銀稽核
 */
type BankDepositsAudit struct {
	Txnid string `json:"txnid"`
}

/**
 * Api Url : bankDeposits/unlock (PUT) 網銀入款列表頁解鎖
 */
type BankDepositsUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}