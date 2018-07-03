package lophiiformesdto

/**
 * API URL : company_bank
 */
type CompanyBankList struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
}

/**
 * API URL : company_bank/data
 */
type CompanyBankData struct {
	Hall_Code string `json:"hall_code"`
	Id string `json:"id"`
}

/**
 * API URL : company_bank (POST)
 */
type CreateCompanyBank struct {
	Hall_Code string `json:"hall_code"`
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
type UpdataCompanyBank struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
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
 * API URL : company_bank (DELETE)
 */
type DeleteCompanyBank struct {
	Hall_Code string `json:"hall_code"`
	Id int `json:"id"`
}

/**
 * Api Url : company_bank/dropdownlist (GET)取得公司銀行卡下拉清單
 */
type CompanyBankDropdownList struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
}