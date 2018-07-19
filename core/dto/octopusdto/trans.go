package octopusdto

/**
 * Api Url : trans/unlock 列表頁解鎖
 */
type TransUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}

/**
 * Api Url : trans/selfunlock 稽核頁解鎖
 */
type TransSelfUnlock struct {
	Txnid string `json:"txnid"`
	Type int `json:"type"`
}