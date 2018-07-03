package lophiiformesdto

/**
 * Api Url : payment-config/list
 */
type PaymentConfigList struct {
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : payment-config (POST)
 */
type CreatePaymentConfig struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Default_Id int `json:"default_id"`
	Front_Name string `json:"front_name"`
	Back_Name string `json:"back_name"`
	Merchant_Id string `json:"merchant_id"`
	Merchant_Secret string `json:"merchant_secret"`
	Extra string `json:"extra"`
	Status int `json:"status"`
}

/**
 * Api Url : payment-config (PUT)
 */
type UpdataPaymentConfig struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
	Site_Code string `json:"site_code"`
	Default_Id int `json:"default_id"`
	Front_Name string `json:"front_name"`
	Back_Name string `json:"back_name"`
	Merchant_Id string `json:"merchant_id"`
	Merchant_Secret string `json:"merchant_secret"`
	Extra string `json:"extra"`
	Status int `json:"status"`
}

/**
 * Api Url : payment-config/info
 */
type PaymentConfigInfo struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * Api Url : payment-config (DELETE)
 */
type DeletePaymentConfig struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
}

/**
 * Api Url : payment_config/default 預設第四方支付列表
 */
type PaymentConfigDefault struct {
	Hall_Code string `json:"hall_code"`
}

/**
 * Api Url : payment-config/banks (GET)
 */
type PaymentConfigBanksInfo struct {
	Hall_Code string `json:"hall_code"`
	Fourth_Id string `json:"fourth_id"`
}

/**
 * Api Url : payment-config/banks (PUT)
 */
type UpdataPaymentConfigBanks struct {
	Hall_Code string `json:"hall_code"`
	Fourth_Id int `json:"fourth_id"`
	Bank_Ids string `json:"bank_ids"`
}