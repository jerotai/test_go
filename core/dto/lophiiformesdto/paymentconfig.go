package lophiiformesdto

/**
 * Api Url : payment-config/list (GET) 取得第四方支付清單
 */
type PaymentConfigList struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : payment-config (POST)
 */
type CreatePaymentConfig struct {
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
type UpdatePaymentConfig struct {
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
	Id string `json:"id"`
}

/**
 * Api Url : payment-config (DELETE)
 */
type DeletePaymentConfig struct {
	Id int `json:"id"`
}

/**
 * Api Url : payment_config/default 預設第四方支付列表
 */
type PaymentConfigDefault struct {
}

/**
 * Api Url : payment-config/banks (GET)
 */
type PaymentConfigBanksInfo struct {
	Fourth_Id string `json:"fourth_id"`
}

/**
 * Api Url : payment-config/banks (PUT)
 */
type UpdatePaymentConfigBanks struct {
	Fourth_Id int `json:"fourth_id"`
	Bank_Ids string `json:"bank_ids"`
}