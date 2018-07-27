package lobsterdto

/**
 * Api Url : report/out_in (GET) 出入款報表
 */
type ReportOotInList struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	Order string `json:"order"`
}

/**
 * Api Url : report/deposit (GET) 入款紀錄表
 */
type ReportDepositList struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	Trans_Id string `json:"trans_id"`
	User string `json:"user"`
	Fourth_Id string `json:"fourth_id"`
	Bank_Id string `json:"bank_id"`
	Company_Bank_Id string `json:"company_bank_id"`
	Type string `json:"type"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : report/withdraw (GET) 出款紀錄表
 */
type ReportWithdrawList struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	Trans_Id string `json:"trans_id"`
	User string `json:"user"`
	Type string `json:"type"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 *  Api Url : report/user_trans (GET) 會員金流
 */
type ReportUserTrans struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
	User string `json:"user"`
	Kind string `json:"kind"`
	Type string `json:"type"`
	Page string `json:"page"`
	Count string `json:"count"`
	Order string `json:"order"`
}