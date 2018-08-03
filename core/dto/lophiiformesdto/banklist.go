package lophiiformesdto

/**
 * API URL : bank_list/transfer
 */
type BankListTransfer struct {
}

/**
 * API URL : bank_list/third
 */
type BankListTthird struct {
}

/**
 * Api Url : bank_list/fourth_third (GET) 取得四方預設三方銀行清單
 */
type BankListFourthThird struct {
	Code string `json:"code"`
}

/**
 * Api Url : bank_list/fourth_transfer (GET) 取得四方預設轉帳銀行清單
 */
type BankListfourthTransfer struct {
	Code string `json:"code"`
}

/**
 * Api Url : bank_list/site (GET) 取得出款銀行設定
 */
type BankListSite struct {
	Status string `json:"status"`
}

/**
 * Api Url : bank_list/backend/site (GET) 後台取得出款銀行設定
 */
type BankListbackendSite struct {
	Site_Code string `json:"site_code"`
	Status string `json:"status"`
}

/**
 * Api Url : bank_list/backend/site (PUT) 後台更新出款銀行設定
 */
type UpdateBankListbackendSite struct {
	Site_Code string `json:"site_code"`
	Banks string `json:"banks"`
}