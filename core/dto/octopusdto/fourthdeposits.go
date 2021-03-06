package octopusdto

/**
 * Api Url : fourthDeposits
 */
type FourthDepositsList struct {
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
 * Api Url : fourthDeposits 四方新增
 */
type CreateFourthDeposits struct {
	Fourth_Id int `json:"fourth_id"`
	Bank_Id int `json:"bank_id"`
	Amount float64 `json:"amount"`
}

/**
 * Api Url : fourthDeposits/limit
 */
type FourthDepositsLimit struct {
}

/**
 * Api Url : fourthDeposits/menu/thirds
 */
type FourthDepositsMenuThirds struct {
}

/**
 * Api Url : fourthDeposits/menu/fourths
 */
type FourthDepositsMenuFourths struct {
	Bank_Id string `json:"bank_id"`
}

/**
 * Api Url : fourthDeposits/menu 四方下拉式選單
 */
type FourthDepositsMenu struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : fourthDeposits/reject 四方拒絕
 */
type FourthDepositsReject struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : fourthDeposits/audit 四方稽核
 */
type FourthDepositsAudit struct {
	Txnid string `json:"txnid"`
}

/**
 * Api Url : fourthDeposits/pass 四方通過
 */
type FourthDepositsPass struct {
	Txnid string `json:"txnid"`
	Front_Remark string `json:"front_remark"`
	Back_Remark string `json:"back_remark"`
}

/**
 * Api Url : fourthDeposits/unlock (PUT) 四方入款列表頁解鎖
 */
type FourthDepositsUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}