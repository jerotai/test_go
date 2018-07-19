package octopusdto

/**
 * Api Url : transReport 入款總匯列表
 */
type TransReportList struct {
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}

/**
 * Api Url : transReports/deposit 存款紀錄
 */
type TransReportDeposit struct {
}

/**
 * Api Url : transReports/withdraw 取款紀錄
 */
type TransReportWithdraw struct {
}