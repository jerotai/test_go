package lobsterdto

/**
 * Api Url : deposit_type (GET) 存款類別
 */
type DepositType struct {
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : withdraw_type (GET) 出款類別
 */
type WithdrawType struct {
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : kind/trans_type (GET) 金流類別
 */
type TransType struct {
	Hall_Code string `json:"hall_code"`
	Kind int `json:"kind"`
}