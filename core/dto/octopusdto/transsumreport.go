package octopusdto

/**
 * Api Url : transSumReport/deposit (GET) 入款總匯列表
 */
type TransSumReportDepositList struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
}


/**
 * Api Url : transSumReport/withdraw (GET) 出款總匯列表
 */
type TransSumReportWithdrawList struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Site_Code string `json:"site_code"`
}