package octopusdto

/**
 * Api Url : trans/unlock 列表頁解鎖
 */
type TransUnlock struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}

/**
 * Api Url : trans/selfunlock 稽核頁解鎖
 */
type TransSelfUnlock struct {
	Hall_Code string `json:"hall_code"`
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}