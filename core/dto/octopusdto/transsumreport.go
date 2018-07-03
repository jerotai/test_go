package octopusdto

/**
 * Api Url : transSumReport/deposit (GET) 入款總匯列表
 */
type TransSumReportDepositList struct {
	Hall_Code string `json:"hall_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}


/**
 * Api Url : transSumReport/withdraw (GET) 出款總匯列表
 */
type TransSumReportWithdrawList struct {
	Hall_Code string `json:"hall_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}