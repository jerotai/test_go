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