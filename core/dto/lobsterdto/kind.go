package lobsterdto

/**
 * Api Url : deposit_type (GET) 存款類別
 */
type DepositType struct {
}

/**
 * Api Url : withdraw_type (GET) 出款類別
 */
type WithdrawType struct {
}

/**
 * Api Url : kind/trans_type (GET) 金流類別
 */
type TransType struct {
	Kind int `json:"kind"`
}