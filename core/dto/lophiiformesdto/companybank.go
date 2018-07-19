package lophiiformesdto

/**
 * API URL : company_bank
 */
type CompanyBankList struct {
	Site_Code string `json:"site_code"`
}

/**
 * API URL : company_bank/data
 */
type CompanyBankData struct {
	Id string `json:"id"`
}

/**
 * API URL : company_bank (POST)
 */
type CreateCompanyBank struct {
	Site_Code string `json:"site_code"`
	Bank_Id int `json:"bank_id"`
	Front_Nickname string `json:"front_nickname"`
	Back_Nickname string `json:"back_nickname"`
	Name string `json:"name"`
	Account string `json:"account"`
	Account_Point string `json:"account_point"`
	Status int `json:"status"`
	Remark string `json:"remark"`
}

/**
 * API URL : company_bank (PUT)
 */
type UpdateCompanyBank struct {
	Id int `json:"id"`
	Bank_Id *int `json:"bank_id"` //使用 *int 表示當前端未送參數時 將過濾 不加入傳送參數 (int 空值時 預設為0)
	Front_Nickname string `json:"front_nickname"`
	Back_Nickname string `json:"back_nickname"`
	Name string `json:"name"`
	Account string `json:"account"`
	Account_Point string `json:"account_point"`
	Status int `json:"status"`
	Remark string `json:"remark"`
}

/**
 * API URL : company_bank (DELETE)
 */
type DeleteCompanyBank struct {
	Id int `json:"id"`
}

/**
 * Api Url : company_bank/dropdownlist (GET)取得公司銀行卡下拉清單
 */
type CompanyBankDropdownList struct {
	Site_Code string `json:"site_code"`
}