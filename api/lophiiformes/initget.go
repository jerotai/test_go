package lophiiformes

import (
	dto "Stingray/core/dto/lophiiformesdto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (l *Lophiiformes) InitGetApiConfig(apiUrl string) (string, interface{}){
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/transfer_bank/list":
		apiDto = &dto.BankListTransfer{}
		apiRequestUrl = "bank_list/transfer"
	case "/third_bank/list":
		apiDto = &dto.BankListTthird{}
		apiRequestUrl = "bank_list/third"
	case "/company_bank/list":
		apiDto = &dto.CompanyBankList{}
		apiRequestUrl = "company_bank"
	case "/company_bank/data":
		apiDto = &dto.CompanyBankData{}
		apiRequestUrl = "company_bank/data"
	case "/company_bank/dropdownList":
		apiDto = &dto.CompanyBankDropdownList{}
		apiRequestUrl = "company_bank/dropdownlist"
	case "/payment_config/list":
		apiDto = &dto.PaymentConfigList{}
		apiRequestUrl = "payment_config/list"
	case "/payment_config/info":
		apiDto = &dto.PaymentConfigInfo{}
		apiRequestUrl = "payment_config/info"
	case "/payment_config/banks":
		apiDto = &dto.PaymentConfigBanksInfo{}
		apiRequestUrl = "payment_config/banks"
	case "/payment_config/default":
		apiDto = &dto.PaymentConfigDefault{}
		apiRequestUrl = "payment_config/default"
	case "/user_bank/list":
		apiDto = &dto.UserBankList{}
		apiRequestUrl = "user_bank"
	case "/user_bank/info":
		apiDto = &dto.UserBankInfo{}
		apiRequestUrl = "user_bank/info"
	case "/user_bank/backendList":
		apiDto = &dto.BackEndUserBankList{}
		apiRequestUrl = "user_bank/backend/list"
	case "/user_bank/backendInfo":
		apiDto = &dto.BackEndUserBankInfo{}
		apiRequestUrl = "user_bank/backend"
	default:
		helper.HelperLog.ErrorLog("[Lophiiformes InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
		//todo
	}
	
	return apiRequestUrl, apiDto
}
