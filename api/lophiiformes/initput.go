package lophiiformes

import (
	dto "Stingray/core/dto/lophiiformesdto"
	
	"Stingray/helper"
)

/**
 * init Api Put Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (l *Lophiiformes) InitPutApiConfig(apiUrl string) (string, interface{})  {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/company_bank":
		apiDto = &dto.UpdateCompanyBank{}
		apiRequestUrl = "company_bank"
	case "/payment_config":
		apiDto = &dto.UpdatePaymentConfig{}
		apiRequestUrl = "payment_config"
	case "/payment_config/banks":
		apiDto = &dto.UpdatePaymentConfigBanks{}
		apiRequestUrl = "payment_config/banks"
	case "/user_bank":
		apiDto = &dto.UpdateUserBank{}
		apiRequestUrl = "user_bank"
	case "/user_bank/backend":
		apiDto = &dto.BackEndUpdateUserBank{}
		apiRequestUrl = "user_bank/backend"
	case "/bank_list/backendSite":
		apiDto = &dto.UpdateBankListbackendSite{}
		apiRequestUrl = "bank_list/backend/site"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lophiiformes InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}